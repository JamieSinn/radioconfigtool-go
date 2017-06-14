package netconfig

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"firstinspires.org/radioconfigtool/util"
	"bytes"
	"firstinspires.org/radioconfigtool/imaging"
)

// ReadARP watches a handle for incoming ARP requests from the OpenMesh radios.
// ReadARP loops until the Destination Hardware Address is not an empty string.
// NOTE: This is threadblocking, and should be.
func readARP() *layers.ARP {
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
			request := arpLayer.(*layers.ARP)
			if request.Operation == layers.ARPRequest && bytes.Compare(request.SourceProtAddress, []byte{192, 168, 100, 9}) == 0 {
				util.Debug("Got request request")
				util.Debug(request)
				util.Debug(string(request.DstHwAddress))

				if bytes.Compare(request.DstHwAddress, []byte{0, 0, 0, 0, 0, 0}) != 0 {
					return request
				}
			}
		}
	}
}


func WaitForRadioModel() string {
	arpReq := readARP()
	return string(arpReq.DstHwAddress)
}
