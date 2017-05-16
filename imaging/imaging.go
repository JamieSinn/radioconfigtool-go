package imaging

import (
	"os/exec"
	"firstinspires.org/radioconfigtool/netconfig"
	"strconv"
)

type Radio struct {
	Team      int
	SSID      string
	WPAKey    string
	Country   string
	DefaultIP string
	EventData EventConfig
}

type EventConfig struct {
	PracticeFieldID int
	PracticeRadioID int
	Firewall        bool
	BWLimit         bool
}

func (r Radio) Event_init() {
	r.EventData.Firewall = true
	r.EventData.BWLimit = true
}

func (r Radio) OM5P_AN_init() {
	r.DefaultIP = "192.168.1.1"
}

func (r Radio) OMP5_AC_init() {
	r.DefaultIP = "192.168.1.1"
}

func (r Radio) OPM5_Flash() {
	exec.Command("ap51-flash.exe", netconfig.NETINT_LAN, strconv.Itoa(r.Team)+"-AN.bin", strconv.Itoa(r.Team)+"-AC.bin")
}

func (r Radio) Reset() {

}

func (r Radio) Ping() {

}

func (r Radio) CompetitionFlash() {

}

func (r Radio) PracticeFieldFlash() {

}
