package util

import "fmt"

const (
	debug = true
)

func Debug(message ...interface{}) {
	if debug {
		fmt.Print("[DEBUG]")
		fmt.Println(message)
	}
}
