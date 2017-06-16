package util

import (
	"fmt"
	"firstinspires.org/radioconfigtool/config"
	"strconv"
)

// DEBUG messages and interfaces. Only prints if debug=true.
func Debug(message ...interface{}) {
	if config.DEBUG {
		fmt.Print("[DEBUG]")
		fmt.Println(message)
	}
}

func BoolToStr(b bool) string {
	if b {
		return "1"
	}
	return "0"
}

func IsValidTeamNumber(team string) bool {
	num, err := strconv.Atoi(team)
	if err != nil {
		return false
	}
	return validTeamNumber(num)
}

func validTeamNumber(team int) bool {
	return team >= 1 && team <= 9999
}
