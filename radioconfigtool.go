package main

import (
	"firstinspires.org/radioconfigtool/netconfig"
	"firstinspires.org/radioconfigtool/gui"
	"firstinspires.org/radioconfigtool/config"
	. "firstinspires.org/radioconfigtool/imaging"
	"firstinspires.org/radioconfigtool/fileio"
	"bytes"
	"io"
	"time"
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

	if config.EventMode() {
		fileio.LoadKeys()
	}

	_, err := gui.DrawGUI(config.EventMode(), Competition, Home)
	if err != nil {
		panic(err)
	}

	netconfig.GetNETINT_LAN_GUID()

	OM5P_AC.VerifyImage()
	OM5P_AN.VerifyImage()

}

func Home(flash bool, team, wpakey string) {

	/*
	X- Tool opens
	X- Team enters their number
	X- Instructions are on the page
	X- Selects either program, or image buttons.
	- On selecting the program button, it sends the configuration string to the team. (pending changes to the config string and system) Return to main screen.
	- On selecting the image button, listens for ARP request, get radio model, flash radio model via tftp. Return to main screen.
	*/
	if flash {
		// Listen for ARP
		// Start TFTP Server
		// Return once all files are requested.
		// Popup saying complete...
		//TODO: Flash radio
	} else {
		configuration := RouterConfiguration{
			// Compat for 2.4 networks, create a 2.4 and 5ghz network.
			Mode:        "AP25",
			Team:        team,
			WPAKey:      wpakey,
			SSID:        team,
			Firewall:    false,
			BWLimit:     true,
			DHCPEnabled: true,
			RadioID_24:  0,
			RadioID_5:   0,
			Event:       "",
		}
		str := configuration.BuildConfigString()
		enc := EncryptConfigString(str)
		err := SendConfiguration(enc)

		if err != nil {
			switch err.Error() {
			case "Invalid":
				gui.InvalidResp()
				return
			case "OutOfDate":
				gui.OutOfDate()
				return
			case "AtEvent":
				gui.ErrorBox("Error", "Your radio was last programmed at an event, and the event's expiry is in the future. " +
					"To prevent connection issues, please wait until the event is over.")
				return
			}
		}
	}
}

func Competition(team string) {

	/*
	X- Tool opens
	X- Team enters their number
	X- Instructions are on the page
	X- Selects "Program"
	X- Listens for ARP string and gets model
	X- Flashes radio with image
	- Upon the radio booting up again, the radio is configured
	*/

	// Listen for ARP
	// Start TFTP Server
	// Return once all files are requested.
	// Wait for radio to come back up
	// Configure

	configuration := RouterConfiguration{
		Mode:        "B5",
		Team:        team,
		WPAKey:      fileio.GetTeamKey(team),
		SSID:        team,
		Firewall:    true,
		BWLimit:     true,
		DHCPEnabled: false,
		RadioID_24:  0,
		RadioID_5:   0,
		Event:       fileio.GetTeamKey("EVENTEND"),
	}
	str := configuration.BuildConfigString()
	enc := EncryptConfigString(str)
	err := SendConfiguration(enc)

	if err != nil {
		switch err.Error() {
		case "Invalid":
			gui.InvalidResp()
			return
		case "OutOfDate":
			//TODO: Flash radio.
			break
		}
	}
}

// Experimental flashing system
func flash() {
	// NOTE: BLOCKING THREAD
	model := netconfig.WaitForRadioModel()
	radio := RobotRouter{}
	for _, r := range validRouters {
		if model == r.ARPString {
			radio = r
		}
	}
	go netconfig.StartTFTPServer(radio.ReadHandler)
}


