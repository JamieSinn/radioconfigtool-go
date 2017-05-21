package gui

import (
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/walk"
)

var (
	teamNumber, wpaKey *walk.LineEdit
	mainWindow *walk.MainWindow
)

func DrawGUI(event bool) (int, error) {

	return MainWindow{
		Title:   "FRC Radio Configuration Utility",
		MinSize: Size{1280, 720},
		Layout:  Grid{},
		AssignTo: &mainWindow,
		Background: SolidColorBrush{
			Color: 0xDFDFDC,
		},
		Children: []Widget{
			getUserInputs(event),

			VSpacer{
				Row:        1,
				Column:     0,
				ColumnSpan: 3,
				Size:       2,
			},

			getConfigureButton(),

			Composite{
				Column:     0,
				Row:        3,
				ColumnSpan: 3,
				Layout:     Grid{},
				Children: []Widget{
					// Left hand side instructions
					getProgrammingInstructions(),
					// Right hand side instructions
					getResetInstructions(),
				},
			},

		},
	}.Run()
}

func getConfigureButton() Widget {
	return PushButton{
		Column:     1,
		ColumnSpan: 1,
		Row:        2,
		Text:       "Configure",
		OnClicked: func() {
			walk.MsgBox(mainWindow, "Configuring...", "Configuring the radio for team " + teamNumber.Text() + "...", walk.MsgBoxIconInformation)

			//TODO: Handle configuration
		},
		//Image:
		Font: tahoma(30, true),
	}
}

func getUserInputs(event bool) Widget {
	return Composite{
		Column:     0,
		Row:        0,
		ColumnSpan: 3,
		Layout:     Grid{},
		Children: []Widget{
			// First Row
			Label{
				Row:    0,
				Column: 0,
				Text:   "Enter Team Number:",
				Font:   tahoma(50, true),
			},
			LineEdit{
				Row:        0,
				Column:     1,
				ColumnSpan: 2,
				Name:       "Team Number",
				AssignTo: &teamNumber,
				ToolTipText:        "Enter team number",
				MaxLength:          4,
				AlwaysConsumeSpace: true,
				MinSize:            Size{Width: 250, Height: 25},
				MaxSize:            Size{Width: 250, Height: 25},

				Font: tahoma(72, false),
			},

			Label{
				Row:    1,
				Column: 0,
				Text:   "WPA Key:",
				Font:   tahoma(50, true),
				Visible: !event,
			},
			LineEdit{
				Row:        1,
				Column:     1,
				ColumnSpan: 2,
				AssignTo: &wpaKey,
				ToolTipText:        "Enter WPA Key",
				MaxLength:          4,
				AlwaysConsumeSpace: true,
				MinSize:            Size{Width: 250, Height: 25},
				MaxSize:            Size{Width: 250, Height: 25},

				Font:     tahoma(72, false),
				Visible: !event,
			},
		},

	}
}

func getProgrammingInstructions() Widget {
	return Composite{
		Column:     0,
		Row:        0,
		ColumnSpan: 1,
		Layout:     Grid{},
		Children: []Widget{
			Label{
				Row:    0,
				Column: 0,
				Text:   "To program your wireless bridge:",
				Font:   tahoma(12, true),
			},
			Label{
				Row:    1,
				Column: 0,
				Text:   "1) Connect power and Ethernet to the Wireless Bridge.",
				Font:   tahoma(12, false),
			},
			Label{
				Row:    2,
				Column: 0,
				Text:   "2) Make sure to use the \"802.3af\" Ethernet port as shown above.",
				Font:   tahoma(12, false),
			},
			Label{
				Row:    3,
				Column: 0,
				Text:   "3) Wait for the Power light to turn and stay solid.",
				Font:   tahoma(12, false),
			},
			Label{
				Row:    4,
				Column: 0,
				Text:   "4) Enter your team number, and a WPA key (optional), above.",
				Font:   tahoma(12, false),
			},
			Label{
				Row:    5,
				Column: 0,
				Text:   "5) Press \"Configure\", the process should take 15-60 seconds.",
				Font:   tahoma(12, false),
			},
		},
	}
}

func getResetInstructions() Widget {
	return Composite{
		Column:     1,
		Row:        0,
		ColumnSpan: 1,
		Layout:     Grid{},
		Children: []Widget{
			Label{
				Row:    0,
				Column: 0,
				Text:   "If asked to reset your wireless bridge:",
				Font:   tahoma(12, true),
			},
			Label{
				Row:    1,
				Column: 0,
				Text:   "1) Connect power and Ethernet to the Wireless Bridge.",
				Font:   tahoma(12, false),
			},
			Label{
				Row:    2,
				Column: 0,
				Text:   "2) Make sure to use the \"802.3af\" Ethernet port as shown above.",
				Font:   tahoma(12, false),
			},
			Label{
				Row:    3,
				Column: 0,
				Text:   "3) Wait for the Power light to turn and stay solid.",
				Font:   tahoma(12, false),
			},
			Label{
				Row:    4,
				Column: 0,
				Text:   "4) Enter your team number, and a WPA key (optional), above.",
				Font:   tahoma(12, false),
			},
			Label{
				Row:    5,
				Column: 0,
				Text:   "5) Press \"Configure\", the process should take 15-60 seconds.",
				Font:   tahoma(12, false),
			},
		},
	}
}

func tahoma(pointsize int, bold bool) Font {
	return Font{
		Family:    "Tahoma",
		Bold:      bold,
		PointSize: pointsize,
	}
}

