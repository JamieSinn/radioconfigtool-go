package imaging

import (
	"firstinspires.org/radioconfigtool/imaging/fileio"
	"io/ioutil"
)

var (
	OM5P_AC = RobotRadio{
		Image: fileio.RouterImage{
			Path: "firmwareOM5PAC.bin",
			ExpectedSize: 3180218,
		},
		Model: "OM5P-AC",
	}
	OM5P_AN = RobotRadio{
		Image: fileio.RouterImage{
			Path: "firmwareOM5PAN.bin",
			ExpectedSize: 3180361,
		},
		Model: "OM5P-AN",
	}
)

type RobotRadio struct {
	Image fileio.RouterImage
	Model string
}

func (radio RobotRadio) VerifyImage() {
	data, err := ioutil.ReadFile(radio.Image.Path)

	if err != nil {
		panic("Could not find image for " + radio.Model + " at " + radio.Image.Path)
	}

	if fileio.VerifyImage(data, radio.Image, radio.Image.ExpectedSize) {
		panic("Could not verify image for " + radio.Model + " at " + radio.Image.Path)
	}
}