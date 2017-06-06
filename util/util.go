package util

import (
	"fmt"
	"firstinspires.org/radioconfigtool/config"
)

// DEBUG messages and interfaces. Only prints if debug=true.
func Debug(message ...interface{}) {
	if config.DEBUG {
		fmt.Print("[DEBUG]")
		fmt.Println(message)
	}
}

// Logs out to the logging file. Only used in competition mode.
func Log(message ...interface{}) {
	if config.EventMode {

	}
}


func BoolToStr(b bool) string {
	if b {
		return "1"
	}
	return "0"
}