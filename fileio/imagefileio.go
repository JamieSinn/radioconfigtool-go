package fileio

import (
	"fmt"
	"strconv"
	"strings"
	"errors"

	"firstinspires.org/radioconfigtool/util"
)

const (
	FILE_OFFSET = 0x10000

	FWUPGRADECFG     = "fwupgrade.cfg"
	FWUPGRADECFG_SIG = "fwupgrade.cfg.sig"
)

// RouterImage is the type to contain the information that's needed to image a radio using the contained image.
// It holds anything needed by the custom flashing tool.
type RouterImage struct {
	Description  string
	Path         string
	Size         int
	ExpectedSize int
	Files        []RouterImageFile
}

// GetFile returns the RouterImageFile for the file contained within the combined ext image.
func (router RouterImage) GetFile(filename string) (RouterImageFile, error) {
	for _, f := range router.Files {
		if f.Name == filename {
			return f, nil
		}
	}
	return RouterImageFile{}, errors.New("File Not Found")
}

// RouterImageFile is a data holder for the individual files in the ce image.
type RouterImageFile struct {
	Offset int
	Data   []byte
	Size   int
	Name   string
}

// String() prints out the file's identifying information for debugging
func (file RouterImageFile) String() string {
	return file.Name + ": Offset:" + strconv.Itoa(file.Offset) + " Size:" + strconv.Itoa(file.Size)
}

// Translated version of ap51-flash's ce_verify
// Returns true if the image is valid, returns false otherwise.
// A valid image has all the files that are listed in the fwupgrade.cfg file
func VerifyImage(data []byte, router RouterImage, expectedSize int) bool {

	// return value
	ret := 0
	// Version of the combined ext router (formatting setup)
	ce_version := 0
	// Number of files in the combined ext router
	num_files := 0
	// Router name
	router_name := ""
	// Length of file
	data_len := len(data)
	// Filename start
	hdr_offset := 0
	// File MD5 start??
	hdr_offset_sec := 0
	// Total size of the router
	image_size := 0
	// Used for multiple fwupgrade.cfg files - only add the smaller one to the total router size as the other is only a signature.
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

	// Get version of the combined ext router - we only support version 1
	ret, err := fmt.Sscanf(string(data), "CE%02x", &ce_version)

	if err != nil {
		return false
	}

	// Different format for different versions
	switch ce_version {
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

	// Loop over all files that were listed in fwupgrade.cfg
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
		case 1:
			// Starting from num_files offset +1
			// 32 len string (filename) - 8 len int zero padded (file size) - 32 len string (file md5)
			// buff = 0?
			ret, err = fmt.Sscanf(string(data[hdr_offset:]), "%32s%08x%32s", &file_name, &file_size, &file_md5)

			// Make sure all three data points were found
			if ret != 3 {
				return false
			}

			break
		default:
			// Unsupported
			return false
		}
		// Create a type to hold the information of this file
		image_file := RouterImageFile{
			Name:   file_name,
			Size:   file_size,
			Offset: file_offset,
			// Take the data from the beginning of the offset, to the current offset plus the file size
			Data:   data[file_offset:file_offset+file_size],
		}

		// Debug out the current file's info
		util.Debug(image_file.String())

		// Add this file to the master list
		router.Files = append(router.Files, image_file)

		// Shift the offset up to the next file
		file_offset += file_size
		// Shift the hdr offset up
		hdr_offset += hdr_offset_sec
		// Take one down...
		num_files--

		if strings.HasPrefix(file_name, FWUPGRADECFG) {
			// Check if the filename is fwupgrade.cfg and not the .sig file
			if len(FWUPGRADECFG)+1 < len(file_name) && !strings.HasSuffix(file_name, ".sig") {
				// Set the description for the file
				description := data[len(FWUPGRADECFG)+1:]
				router.Description = string(description)
			}

			/***
			 * In case this CE router contains multiple fwupgrade.cfg entries
			 * only the smaller fwupgrade.cfg should be added to the total
			 * router size in order to detect the end-of-flash correctly.
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
	if image_size > expectedSize {
		fmt.Println("Error - bogus CE router: claimed size bigger than actual size: " + strconv.Itoa(image_size))
		return false
	}
	router.Size = image_size
	return true
}
