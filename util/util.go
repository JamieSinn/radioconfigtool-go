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
