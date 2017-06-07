package main

import (
	"firstinspires.org/radioconfigtool/netconfig"
	"firstinspires.org/radioconfigtool/gui"
	"firstinspires.org/radioconfigtool/config"
	. "firstinspires.org/radioconfigtool/imaging"
	"firstinspires.org/radioconfigtool/fileio"
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
	validRouters = RobotRadioList{OM5P_AN, OM5P_AC}
)

func main() {

	_, err := gui.DrawGUI(config.EventMode())
	if err != nil {
		panic(err)
	}

	netconfig.GetNETINT_LAN_GUID()

	OM5P_AC.VerifyImage()
	OM5P_AN.VerifyImage()
}

func TeamFlow() {
	/*
	X- Tool opens
	- Team enters their number
	- Instructions are on the page
	- Selects either program, or image buttons.
	- On selecting the program button, it sends the configuration string to the team. (pending changes to the config string and system) Return to main screen.
	- On selecting the image button, listens for ARP request, get radio model, flash radio model via tftp. Return to main screen.
	 */
}


func DetectRadio() RobotRouter {

	arpresp := netconfig.ReadARP()

	for _, router := range validRouters {
		if bytes.Equal(arpresp.SourceProtAddress, router.ConfigIP) &&
			string(arpresp.DstHwAddress) == router.ARPString {
			return router
		}
	}
	return RobotRouter{}
}