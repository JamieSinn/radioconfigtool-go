package imaging

import (
	"firstinspires.org/radioconfigtool/imaging/fileio"
	"io/ioutil"
	"firstinspires.org/radioconfigtool/netconfig"
	"bytes"
)

var (
	OM5P_AC = RobotRadio{
		Image: fileio.RouterImage{
			Path:         "firmwareOM5PAC.bin",
			ExpectedSize: 3180218,
		},
		Model:     "OM5P-AC",
		ARPString: "OM5PAC",
		ConfigIP:  []byte{192, 168, 100, 9},
	}
	OM5P_AN = RobotRadio{
		Image: fileio.RouterImage{
			Path:         "firmwareOM5PAN.bin",
			ExpectedSize: 3180361,
		},
		Model:     "OM5P-AN",
		ARPString: "OM5PAN",
		ConfigIP:  []byte{192, 168, 100, 9},
	}
	// Other radio models can be added here for future changes
	radios = RobotRadioList{OM5P_AN, OM5P_AC}
)

type RobotRadio struct {
	Image     fileio.RouterImage
	Model     string
	ARPString string
	ConfigIP  []byte
}

type RobotRadioList []RobotRadio

func (radio RobotRadio) VerifyImage() {
	data, err := ioutil.ReadFile(radio.Image.Path)

	if err != nil {
		panic("Could not find image for " + radio.Model + " at " + radio.Image.Path)
	}

	if fileio.VerifyImage(data, radio.Image, radio.Image.ExpectedSize) {
		panic("Could not verify image for " + radio.Model + " at " + radio.Image.Path)
	}
}

func DetectRadio() RobotRadio {

	arpresp := netconfig.ReadARP()

	for _, radio := range radios {
		if bytes.Equal(arpresp.SourceProtAddress, radio.ConfigIP) &&
			string(arpresp.DstHwAddress) == radio.ARPString {
			return radio
		}
	}
}
