package imaging

import (
	"io"
	"os"
	"fmt"
	"time"
	"github.com/pin/tftp"
)

// OpenMesh looks for 192.168.100.8 for the tftp server/client.


// readHandler is called when client starts file download from server
func readHandler(filename string, rf io.ReaderFrom) error {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return err
	}
	n, err := rf.ReadFrom(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return err
	}
	fmt.Printf("%d bytes sent\n", n)
	return nil
}
func main() {
	// use nil in place of handler to disable read or write operations
	s := tftp.NewServer(readHandler, nil)
	s.SetTimeout(5 * time.Hour) // optional
	err := s.ListenAndServe(":69") // blocks until s.Shutdown() is called
	if err != nil {
		fmt.Fprintf(os.Stdout, "server: %v\n", err)
		os.Exit(1)
	}
}