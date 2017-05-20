package imaging

import (
	"os/exec"
	"firstinspires.org/radioconfigtool/netconfig"
	"strconv"
	"net"
	"io/ioutil"
	"firstinspires.org/radioconfigtool/util"
	"strings"
)

type Radio struct {
	Team           int
	SSID           string
	WPAKey         string
	Country        string
	DefaultIP      string
	DefaultNetwork string
	EventData      EventConfig
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
	r.DefaultNetwork = "192.168.1."
}

func (r Radio) OMP5_AC_init() {
	r.DefaultIP = "192.168.1.1"
	r.DefaultNetwork = "192.168.1."
}

func (r Radio) OPM5_Flash() {
	netconfig.SetNetworkAdapterIP(r.DefaultNetwork+"2", "255.255.255.0", r.DefaultIP)
	exec.Command("ap51-flash.exe", "\""+netconfig.NETINT_LAN_GUID+"\"", strconv.Itoa(r.Team)+"-AN.bin", strconv.Itoa(r.Team)+"-AC.bin")
}

func Ping(addr string) {

}

func CheckImage() bool {
	ret := false
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {

	}
	conn, err := listener.Accept()
	if err != nil {

	}
	go func(conn net.Conn) {
		if conn == nil {
			return
		}
		util.Debug("Connection from: " + conn.RemoteAddr().String())

		result, err := ioutil.ReadAll(conn)
		if err != nil {
			panic(err)
		}
		split := strings.Split(string(result), "\n")

		if strings.HasPrefix(split[0], "Team:") &&
			strings.HasPrefix(split[1], "Event:") {
			ret = true
		}
	}(conn)
	return ret
}
func handleConnection(conn net.Conn) {

}
