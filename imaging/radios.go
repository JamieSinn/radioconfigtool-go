package imaging

import (
	"firstinspires.org/radioconfigtool/imaging/fileio"
	"io/ioutil"
	"firstinspires.org/radioconfigtool/netconfig"
	"bytes"
)

var (
	OM5P_AC = RobotRouter{
		Image: fileio.RouterImage{
			Path:         "firmwareOM5PAC.bin",
			ExpectedSize: 3180218,
		},
		Model:     "OM5P-AC",
		ARPString: "OM5PAC",
		ConfigIP:  []byte{192, 168, 100, 9},
	}
	OM5P_AN = RobotRouter{
		Image: fileio.RouterImage{
			Path:         "firmwareOM5PAN.bin",
			ExpectedSize: 3180361,
		},
		Model:     "OM5P-AN",
		ARPString: "OM5PAN",
		ConfigIP:  []byte{192, 168, 100, 9},
	}
	// Other router models can be added here for future changes
	routers = RobotRadioList{OM5P_AN, OM5P_AC}
)

type RobotRouter struct {
	Image     fileio.RouterImage
	Model     string
	ARPString string
	ConfigIP  []byte
}

type RobotRadioList []RobotRouter

func (router RobotRouter) VerifyImage() {
	data, err := ioutil.ReadFile(router.Image.Path)

	if err != nil {
		panic("Could not find image for " + router.Model + " at " + router.Image.Path)
	}

	if fileio.VerifyImage(data, router.Image, router.Image.ExpectedSize) {
		panic("Could not verify image for " + router.Model + " at " + router.Image.Path)
	}
}

func DetectRadio() RobotRouter {

	arpresp := netconfig.ReadARP()

	for _, router := range routers {
		if bytes.Equal(arpresp.SourceProtAddress, router.ConfigIP) &&
			string(arpresp.DstHwAddress) == router.ARPString {
			return router
		}
	}
	return RobotRouter{}
}
