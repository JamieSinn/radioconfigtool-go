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

// OpenMesh looks for 192.168.100.8 for the tftp server/client.
func TFTPInit() {
	SetNetworkAdapterIP("192.168.100.8", "255.255.255.0", "192.168.100.1")

}

// readHandler is called when client starts file download from server
func readHandler(filename string, rf io.ReaderFrom) error {
	// TODO: have a proper getter per cycle.
	radio := imaging.RobotRouter{}

	file, err := radio.Image.GetFile(filename)
	if err != nil {
		util.Debug("Could not find requested file: " + filename)
		return err
	}
	n, err := rf.ReadFrom(bytes.NewReader(file.Data))
	if err != nil {
		util.Debug("%v", err)
		return err
	}
	util.Debug("%d bytes sent", n)
	return nil
}

// TODO: Clean up initialization, and make it a bit easier to work with to serve the specific files.
func StartTFTPServer() {
	// use nil in place of handler to disable read or write operations
	s := tftp.NewServer(readHandler, nil)
	s.SetTimeout(5 * time.Hour)    // optional
	err := s.ListenAndServe(":69") // blocks until s.Shutdown() is called
	if err != nil {
		fmt.Fprintf(os.Stdout, "server: %v\n", err)
		os.Exit(1)
	}
}
