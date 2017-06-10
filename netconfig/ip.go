package netconfig

import (
	"os/exec"

	"strings"
	"bytes"

	"firstinspires.org/radioconfigtool/config"
	"firstinspires.org/radioconfigtool/util"
	"github.com/tatsushid/go-fastping"
	"net"
	"time"
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

// GetNETINT_LAN_GUID gets the GUID for the LAN interface for use with pcap
func GetNETINT_LAN_GUID() {
	getmac := exec.Command("getmac", "/nh", "/v", "/fo", "csv")
	output, _ := getmac.StdoutPipe()
	getmac.Start()
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(output)
	raw := buffer.String()
	raw = strings.Replace(raw, "\"", "", -1)
	raw = strings.Replace(raw, "\r", ",", -1)
	split := strings.Split(raw, ",")
	for i, s := range split {
		util.Debug("Found network interface: \"" + s + "\"")
		if s == config.NETINT_LAN {
			NETINT_LAN_GUID = strings.Replace(split[i+3], "Tcpip", "NPF", -1)
			return
		}
	}
}

func Ping(ip string) bool {
	ret := false
	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip4:icmp", ip)
	if err != nil {
		util.Debug(err)
		return false
	}

	p.AddIPAddr(ra)
	p.MaxRTT = 5 * time.Second

	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		util.Debug("IP Addr: " + addr.String() + " receive, RTT: " + rtt.String())
		ret = true
	}
	p.OnIdle = func() {
		util.Debug("Finished pinging " + ip)
	}
	err = p.Run()
	if err != nil {
		util.Debug(err)
	}
	return ret
}
