package netconfig

import (
	"io"
	"os"
	"fmt"
	"time"
	"github.com/pin/tftp"
	"bytes"
	"firstinspires.org/radioconfigtool/util"
	"firstinspires.org/radioconfigtool/imaging"
)

var (
	TFTPServer *tftp.Server
)
// OpenMesh looks for 192.168.100.8 for the tftp server/client.
func StartTFTPServer(readHandler func(filename string, rf io.ReaderFrom) error) {
	SetNetworkAdapterIP("192.168.100.8", "255.255.255.0", "")
	// use nil in place of handler to disable read or write operations
	s := tftp.NewServer(readHandler, nil)
	TFTPServer = s
	err := s.ListenAndServe(":69") // blocks until s.Shutdown() is called
	if err != nil {
		fmt.Fprintf(os.Stdout, "server: %v\n", err)
		s.Shutdown()
		return
	}
}
