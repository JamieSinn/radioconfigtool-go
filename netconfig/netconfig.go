package netconfig

import (
	"os/exec"

	"strings"
	"bytes"

	"fmt"
	"firstinspires.org/radioconfigtool/config"
)

var (
	NETINT_LAN_GUID = ""
)

// ResetNetworkAdapter resets the network adapter back to DHCP addressing
func ResetNetworkAdapter(inter string) {
	exec.Command("netsh", "interface", "ipv4", "set", "address", "name=\""+inter+"\"", "dhcp")
	exec.Command("netsh", "interface", "ipv4", "set", "dns", "name=\""+inter+"\"", "dhcp")
}

func SetNetworkAdapterIP(ip, netmask, gateway string) {
	exec.Command("netsh", "interface", "ipv4", "set", "address", "name=\""+config.NETINT_LAN+"\"", "static",
		ip, netmask, gateway)
}

// GetNETINT_LAN_GUID gets the GUID for the LAN interface for use with ap51-flash.
func GetNETINT_LAN_GUID() {
	getmac := exec.Command("getmac", "/nh", "/v", "/fo", "csv")
	output, _ := getmac.StdoutPipe()
	getmac.Start()
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(output)
	raw := buffer.String()
	raw = strings.Replace(raw, "\"", "", -1)
	split := strings.Split(raw, ",")
	for i, s := range split {
		fmt.Println(s)
		if s == config.NETINT_LAN {
			NETINT_LAN_GUID = strings.Replace(split[i+3], "Tcpip", "NPF", -1)
			return
		}
	}
}
