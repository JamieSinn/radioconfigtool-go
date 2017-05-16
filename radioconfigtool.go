package main

import (
	"github.com/lxn/walk"
	"strings"
	"net"
	"firstinspires.org/radioconfigtool/netconfig"
	"fmt"
	"runtime"
	"github.com/lxn/walk/declarative"
)

func main() {

	if runtime.GOOS != "windows" {
		fmt.Println("The configuration tool has only been built for Windows. Please use a different computer.")
		return
	}

	fmt.Println("Network Interfaces:")
	interfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	for _, i := range interfaces {
		fmt.Println(i)

		fmt.Println("----")
	}
	fmt.Println(net.InterfaceByName(netconfig.NETINT_LAN))

	var inTE, outTE *walk.TextEdit

	declarative.MainWindow{
		Title:   "SCREAMO",
		MinSize: declarative.Size{600, 400},
		Layout:  declarative.VBox{},
		Children: []declarative.Widget{
			declarative.HSplitter{
				Children: []declarative.Widget{
					declarative.TextEdit{AssignTo: &inTE},
					declarative.TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			declarative.PushButton{
				Text: "SCREAM",
				OnClicked: func() {
					outTE.SetText(strings.ToUpper(inTE.Text()))
				},
			},
		},
	}.Run()
}
