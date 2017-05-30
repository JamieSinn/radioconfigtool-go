package main

import (
	"fmt"
	"runtime"
	"os"
	"firstinspires.org/radioconfigtool/eventconfig"
	"firstinspires.org/radioconfigtool/util"
	"firstinspires.org/radioconfigtool/resources"
	"firstinspires.org/radioconfigtool/netconfig"
	"firstinspires.org/radioconfigtool/gui"
	"firstinspires.org/radioconfigtool/config"
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
}


// In order to use the ap51-flash.exe utility, and to have only one .exe to distribute
// ap51-flash.exe needs to be copied out from the resources and into the currnet directory.
// This functionality isn't currently used as I may just package it with it during installation.
func writeOutAP51Flash() {
	file, err := os.Create("ap51-flash.exe")
	if err != nil {
		// Could not create file
	}
	defer file.Close()
	data, err := resources.Asset("ap51-flash.exe")

	if err != nil {
		// Could not read from resource
	}
	file.Write(data)
	file.Sync()
}

