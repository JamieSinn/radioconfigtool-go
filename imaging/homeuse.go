package imaging

import (
	"firstinspires.org/radioconfigtool/netconfig"
	"net"
	"io/ioutil"
	"strings"
)

func sendConfiguration() {
	netconfig.SetNetworkAdapterIP("192.168.1.2", "255.255.255.0", "192.168.1.1")
	conn, err := net.Dial("tcp", "192.168.1.1:8888")
	if err != nil {
		// Could not connect
	}
	valid := checkRadioResponse(conn)
	if !valid {
		//Invalid radio response
	}
}

func checkRadioResponse(conn net.Conn) bool {
	ret := false;
	result, err := ioutil.ReadAll(conn)
	if err != nil {
		panic(err)
	}
	split := strings.Split(string(result), "\n")

	if strings.HasPrefix(split[0], "Team:") &&
		strings.HasPrefix(split[1], "Event:") {
		ret = true
	}
	return ret
}

