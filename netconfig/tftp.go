package netconfig

import (
	"io"
	"github.com/pin/tftp"
	"time"
	"firstinspires.org/radioconfigtool/gui"
	"strconv"
	"firstinspires.org/radioconfigtool/config"
)

var (
	SentFiles = make(chan int)
)
// OpenMesh looks for 192.168.100.8 for the tftp server/client.
func StartTFTPServer(readHandler func(filename string, rf io.ReaderFrom) error, numFiles int) {
	SetNetworkAdapterIP("192.168.100.8", "255.255.255.0", "")
	// use nil in place of handler to disable read or write operations
	s := tftp.NewServer(readHandler, nil)
	go s.ListenAndServe("192.168.100.8:69") // blocks until s.Shutdown() is called
	for {
		i := <-SentFiles
		select {
		case i >= numFiles:
			s.Shutdown()
			return
		case <-time.After(config.TFTP_TIMEOUT):
			gui.ErrorBox("Error", "Failed to send all files within the 5 minute timelimit. " + strconv.Itoa(i) +
				" files were sent of the " + strconv.Itoa(numFiles) + " available")
			s.Shutdown()
			return
		}
	}
}
