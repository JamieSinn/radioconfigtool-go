package netconfig

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"firstinspires.org/radioconfigtool/util"
)

// ReadARP watches a handle for incoming ARP requests from the OpenMesh radios.
// ReadARP loops until the Destination Hardware Address is not an empty string.
func ReadARP() *layers.ARP {
	handle, err := pcap.OpenLive(NETINT_LAN_GUID, 65536, true, pcap.BlockForever)

	if err != nil {
		util.Debug(err)
		return nil
	}

	src := gopacket.NewPacketSource(handle, layers.LayerTypeEthernet)
	in := src.Packets()
	for {
		var packet gopacket.Packet
		select {
		case packet = <-in:
			arpLayer := packet.Layer(layers.LayerTypeARP)
			if arpLayer == nil {
				continue
			}
			arp := arpLayer.(*layers.ARP)
			if arp.Operation == layers.ARPRequest {
				util.Debug("Got arp request")
				util.Debug(arp)
				util.Debug(string(arp.DstHwAddress))

				if string(arp.DstHwAddress) != "" {
					return arp
				}
			}
		}
	}
}


