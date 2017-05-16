package netconfig

import (
	"os/exec"

)

const (
	// NETINT_LAN is the interface name to use for imaging/configuration
	NETINT_LAN  = "LAN"
)

// ResetNetworkAdapter resets the network adapter back to DHCP addressing
func ResetNetworkAdapter() {
	exec.Command("netsh", "interface", "ipv4", "set", "address", "name=\"" +NETINT_LAN+ "\"", "dhcp")
	exec.Command("netsh", "interface", "ipv4", "set", "dns", "name=\"" +NETINT_LAN+ "\"", "dhcp")
}

func SetNetworkAdapterIP(ip, netmask, gateway string) {
	exec.Command("netsh", "interface", "ipv4", "set", "address", "name=\"" +NETINT_LAN+ "\"", "static",
	ip, netmask, gateway)
}
