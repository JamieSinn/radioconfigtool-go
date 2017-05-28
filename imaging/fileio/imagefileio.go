package fileio

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	FILE_OFFSET          = 0x10000
	FILE_NAME_MAX_LENGTH = 33
	DESC_MAX_LENGTH      = 30
	OM5PAC_TEST_STRING   = "CE01OM5PAC                          04fwupgrade.cfg                   0000027ffec6a1afa8b1c578000318d34d2d0991kernel                          001084a63c7a7fb0a93f7685f0108941903042fbrootfs                          00200004601110fc3cad623ae6213c259649834cfwupgrade.cfg.sig               000002102b99ad42124beb0a0e9f825f62098a83"

	FWUPGRADECFG     = "fwupgrade.cfg"
	FWUPGRADECFG_SIG = "fwupgrade.cfg.sig"
)

type RouterImage struct {
	RouterType    int
	Description   string
	Path          string
	EmbeddedImage []byte
	Size          int
	Files         []RouterImageFile
}

func (router RouterImage) getFile(filename string) RouterImageFile {
	for _, f := range router.Files {
		if f.Name == filename {
			return f
		}
	}
	return nil
}

type RouterImageFile struct {
	Offset int
	Data   []byte
	Size   int
	Name   string
}

// Translated version of ap51-flash's ce_verify
// Returns true if the image is valid, returns false otherwise.
func VerifyImage(data []byte, image RouterImage) bool {

	ret := 0
	// Version of the combined ext image (formatting setup)
	ce_version := 0
	// Number of files in the combined ext image
	num_files := 0
	// Router name
	router_name := ""
	// Length of file
	data_len := len(data)
	// Filename start
	hdr_offset := 0
	// File MD5 start??
	hdr_offset_sec := 0
	// Total size of the image
	image_size := 0
	// Used for multiple fwupgrade.cfg files - only add the smaller one to the total image size as the other is only a signature.
	fwcfg_size := 0
	// The start of the proper data for files
	file_offset := FILE_OFFSET


	// Check if the length of the data is valid.
	if data_len < 100 {
		return false
	}

	// Check if the file is not a Combined Ext Image
	if data[0] != 'C' || data[1] != 'E' {
		return false
	}

	// Get version of the combined ext image
	ret, err := fmt.Sscanf(string(data), "CE%02x", &ce_version)

	if err != nil {
		return false
	}


	// Different format for different versions
	switch ce_version {
	case 0:
		ret, err = fmt.Sscanf(string(data), "CE%10s%02x", &router_name, &num_files)
		if ret != 2 {
			return false
		}
		hdr_offset = 14
		hdr_offset_sec = 28
		break

	case 1:
		// Replace the version string into the CE prefix
		format := "CE0" + strconv.Itoa(ce_version) + "%32s%02x"

		// Get router name and number of files
		ret, err = fmt.Sscanf(string(data), format, &router_name, &num_files)

		if ret != 2 {
			return false
		}

		hdr_offset = 38
		hdr_offset_sec = 72
		break
	default:
		// Unsupported
		return false

	}


	// Loop over all files that were found.
	for num_files > 0 {
		// File name
		file_name := ""
		// File size in bytes
		file_size := 0
		// File md5 hash
		file_md5 := ""
		// hdr_offset + hdr_offset_sec = kernel start
		if hdr_offset+hdr_offset_sec > data_len {
			fmt.Println("Error - buffer too small to parse CE header")
			return false
		}

		switch ce_version {
		case 0:
			ret, err = fmt.Sscanf(string(data[hdr_offset:]), "%20s%08x", &file_name, &file_size)
			if ret != 2 {
				return false
			}

			break
		case 1:
			// Starting from num_files offset +1
			// 32 len string (filename) - 8 len int (file size) - 32 len string (file md5)
			// buff = 0?
			ret, err = fmt.Sscanf(string(data[hdr_offset:]), "%32s%08x%32s", &file_name, &file_size, &file_md5)
			if ret != 3 {
				return false
			}

			break
		}

		image_file := RouterImageFile{
			Name:   file_name,
			Size:   file_size,
			Offset: file_offset,
		}
		image.Files = append(image.Files, image_file)

		// Shift the offset up to the next file
		file_offset += file_size
		// Shift the hdr offset up
		hdr_offset += hdr_offset_sec
		// Take one down...
		num_files--

		if strings.HasPrefix(file_name, FWUPGRADECFG) {
			// Check if the filename is fwupgrade.cfg and not the .sig file
			if len(FWUPGRADECFG)+1 < len(file_name) && !strings.HasSuffix(file_name, ".sig") {
				description := data[len(FWUPGRADECFG)+1:]
				image.Description = string(description)
			}

			/***
			 * In case this CE image contains multiple fwupgrade.cfg entries
			 * only the smaller fwupgrade.cfg should be added to the total
			 * image size in order to detect the end-of-flash correctly.
			 */
			if (fwcfg_size > 0) && (fwcfg_size <= file_size) {
				continue
			}
			if fwcfg_size > file_size {
				image_size -= fwcfg_size
			}

			fwcfg_size = file_size
		}

		// increase total image size
		image_size += file_size
	}
	image.Size = image_size
	return true
}
