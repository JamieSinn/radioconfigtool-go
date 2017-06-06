package main

import (
	"firstinspires.org/radioconfigtool/util"
	"firstinspires.org/radioconfigtool/netconfig"
	"firstinspires.org/radioconfigtool/gui"
	"firstinspires.org/radioconfigtool/config"
	"firstinspires.org/radioconfigtool/imaging"
)

func main() {

	_, err := gui.DrawGUI(config.EventMode)
	if err != nil {
		panic(err)
	}

	netconfig.GetNETINT_LAN_GUID()

	imaging.OM5P_AC.VerifyImage()
	imaging.OM5P_AN.VerifyImage()
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