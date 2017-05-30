package main

import (
	"firstinspires.org/radioconfigtool/eventconfig"
	"firstinspires.org/radioconfigtool/util"
	"firstinspires.org/radioconfigtool/netconfig"
	"firstinspires.org/radioconfigtool/gui"
	"firstinspires.org/radioconfigtool/config"
	"firstinspires.org/radioconfigtool/imaging"
)

func main() {

	if config.EventMode {
		util.Debug("Downloading all OpenWRT images for teams.")
		teams := eventconfig.GetTeams()
		eventconfig.GetAllImages(teams)
		util.Debug("Starting GUI with event security enabled.")
	}

	if _, err := gui.DrawGUI(config.EventMode); err != nil {
		panic(err)
	}

	netconfig.GetNETINT_LAN_GUID()

	imaging.OM5P_AC.VerifyImage()
	imaging.OM5P_AN.VerifyImage()
}
