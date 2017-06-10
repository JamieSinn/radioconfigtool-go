package imaging

import (
	"testing"
	"firstinspires.org/radioconfigtool/util"
)

func TestRouterConfiguration_BuildConfigString(t *testing.T) {
	configuration := RouterConfiguration{
		// Compat for 2.4 networks, possibly going to use both 2.4 and 5
		Mode:        "AP24",
		Team:        "1334",
		WPAKey:      "",
		SSID:        "1334",
		Firewall:    false,
		BWLimit:     true,
		DHCPEnabled: true,
		RadioID_24:  0,
		RadioID_5:   0,
		Comment:     "Comment",
	}
	str := configuration.BuildConfigString()
	util.Debug(str)
	check := "AP24,1334,1334,,1,0,1,0,0,Comment,\n"
	if str != check {
		t.Fail()
	}
}