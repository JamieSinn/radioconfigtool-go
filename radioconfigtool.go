package main

import (
	"github.com/lxn/walk"
	"strings"
	"fmt"
	"runtime"
	"github.com/lxn/walk/declarative"
	"os"
	"firstinspires.org/radioconfigtool/eventconfig"
	"firstinspires.org/radioconfigtool/util"
	"firstinspires.org/radioconfigtool/resources"
	"firstinspires.org/radioconfigtool/netconfig"
)

var (
	EventMode = false
	PracticeField = false
)
func main() {

	if runtime.GOOS != "windows" {
		fmt.Println("The configuration tool has only been built for Windows. Please use a different computer.")
		return
	}

	if EventMode {
		util.Debug("Downloading all OpenWRT images for teams.")
		teams := eventconfig.GetTeams()
		eventconfig.GetAllImages(teams)
	}

	netconfig.GetNETINT_LAN_GUID()
}

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

