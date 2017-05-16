package gui

import (
	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"
	"strings"
)

//TODO
func buildGUI() {

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