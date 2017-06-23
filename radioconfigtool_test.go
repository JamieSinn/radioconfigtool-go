package main

import (
	"testing"
	"github.com/pin/tftp"
	"os"
	"fmt"
)

func TestCompetition(t *testing.T) {
	c, err := tftp.NewClient("192.168.100.8:69")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	for _, radio := range validRouters {
		for i, file := range radio.Image.Files {
			wt, err := c.Receive(file.Name, "octet")
			if err != nil {
				fmt.Println(err)
				t.Fail()
			}
			// Optionally obtain transfer size before actual data.
			if n, ok := wt.(tftp.IncomingTransfer).Size(); ok {
				fmt.Printf("Transfer size: %d\n", n)
			}

			n, err := wt.WriteTo(os.Stdout)
			fmt.Printf("%d bytes received\n", n)
			fmt.Printf("%d/%d files received", i+1, len(radio.Image.Files))
		}
	}
}
