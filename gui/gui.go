package gui

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"errors"
	"firstinspires.org/radioconfigtool/util"
)

var (
	WPAKey     *walk.LineEdit
	TeamNumber *walk.LineEdit
	Window     *walk.MainWindow
	isEvent    bool
)
// Used for testing the GUI
func main() {
	DrawGUI(true,
		func(team string) {
			walk.MsgBox(Window, "Configuring...", "Configuring the radio for team "+TeamNumber.Text()+"...", walk.MsgBoxIconInformation)
		},
		func(flash bool, team, wpakey string) {
			walk.MsgBox(Window, "Configuring...", "Configuring the radio for team "+TeamNumber.Text()+"...", walk.MsgBoxIconInformation)
		})
}

func DrawGUI(event bool, competition func(team string), home func(flash bool, team, wpakey string)) (int, error) {
	isEvent = event
	icon, err := walk.NewIconFromResourceId(7)
	if err != nil {
		return 0, errors.New("Failed to load icon from resources.")
	}

	/*
		om5p, err := resources.Asset("om5p.png")

		if err != nil {
			return 0, errors.New("Failed to load OM5P Image from resources.")
		}

		jpegdata, err := jpeg.Decode(bytes.NewReader(om5p))
	*/

	if err != nil {
		return 0, errors.New("Failed to load OM5P Image from resources.")
	}

	return MainWindow{
		Title:    "FRC Radio Configuration Utility",
		MinSize:  Size{Width: 1280, Height: 720},
		Layout:   Grid{},
		AssignTo: &Window,
		Icon:     icon,
		Background: SolidColorBrush{
			Color: 0xDFDFDF,
		},
		Children: []Widget{
			getUserInputs(event),

			VSpacer{
				Row:        1,
				Column:     0,
				ColumnSpan: 3,
				Size:       2,
			},

			getConfigureButton(competition, home),

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

func getConfigureButton(competition func(team string), home func(flash bool, team, wpakey string)) Widget {

	if isEvent {

		return PushButton{
			Column:     1,
			ColumnSpan: 1,
			Row:        2,
			Text:       "Configure",
			OnClicked: func() {
				if !util.IsValidTeamNumber(TeamNumber.Text()) {
					invalidTeam()
					return
				}
				competition(TeamNumber.Text())
			},
			//Image:
			Font: tahoma(30, true),
		}
	} else {
		return Composite{
			Column:     0,
			ColumnSpan: 3,
			Row:        2,
			Layout:     Grid{},
			Children: []Widget{
				PushButton{
					Column:     1,
					ColumnSpan: 2,
					Row:        0,
					Text:       "Configure",
					OnClicked: func() {
						if !util.IsValidTeamNumber(TeamNumber.Text()) {
							invalidTeam()
							return
						}
						home(false, TeamNumber.Text(), WPAKey.Text())
					},
					//Image:
					Font: tahoma(30, true),
				},
				PushButton{
					Column:     3,
					ColumnSpan: 1,
					Row:        0,
					Text:       "Flash",
					OnClicked: func() {
						if !util.IsValidTeamNumber(TeamNumber.Text()) {
							invalidTeam()
							return
						}
						home(true, TeamNumber.Text(), WPAKey.Text())
					},
					//Image:
					Font: tahoma(30, true),
				},
			},
		}
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
				Row:                0,
				Column:             1,
				ColumnSpan:         2,
				Name:               "Team Number",
				AssignTo:           &TeamNumber,
				ToolTipText:        "Enter team number",
				MaxLength:          4,
				AlwaysConsumeSpace: true,
				MinSize:            Size{Width: 250, Height: 25},
				MaxSize:            Size{Width: 250, Height: 25},

				Font: tahoma(72, false),
			},

			Label{
				Row:     1,
				Column:  0,
				Text:    "WPA Key:",
				Font:    tahoma(50, true),
				Visible: !event,
			},
			LineEdit{
				Row:                1,
				Column:             1,
				ColumnSpan:         2,
				AssignTo:           &WPAKey,
				ToolTipText:        "Enter WPA Key",
				MaxLength:          4,
				AlwaysConsumeSpace: true,
				MinSize:            Size{Width: 250, Height: 25},
				MaxSize:            Size{Width: 250, Height: 25},

				Font:    tahoma(72, false),
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
		Column:     2,
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

func invalidTeam() {
	ErrorBox( "Error", "Invalid Team Number. Team number must be 1-9999.")
}

func InvalidResp() {
	ErrorBox( "Error", "Invalid Radio Response. Please make sure you are using a supported radio.")
}

func OutOfDate() {
	ErrorBox( "Error", "Radio firmware is out of date, please re-image.")
}

func InfoBox(title, message string) {
	walk.MsgBox(Window, title, message, walk.MsgBoxIconInformation)
}

func ErrorBox(title, message string) {
	walk.MsgBox(Window, title, message, walk.MsgBoxIconError)
}