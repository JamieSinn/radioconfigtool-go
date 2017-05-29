package netconfig

import (
	"firstinspires.org/radioconfigtool/imaging"
	"github.com/mdlayher/arp"
	"net"
	"firstinspires.org/radioconfigtool/config"
)


const (
	PACKET_BUFF_LEN = 2000
)
// This lovely mess of code (that is to be written) sends an ARP request out to radios to identify them and grab their hardware version

func GetRadioHardwareAddress() net.HardwareAddr {
	SetNetworkAdapterIP("192.168.100.8", "255.255.255.0", "192.168.100.1")

	ifi, err := net.InterfaceByName(config.NETINT_LAN)
	if err != nil {

	}
	client, err := arp.Dial(ifi)
	if err != nil {

	}
	addr, err := client.Resolve(net.IP{192, 168, 100, 9})
	if err != nil {

	}
	return addr
}

func GetRadioType() imaging.RobotRadio {

	return imaging.OM5P_AN
}
