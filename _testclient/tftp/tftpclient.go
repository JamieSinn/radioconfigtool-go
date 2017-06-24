package main

import (
	"github.com/pin/tftp"
	"fmt"
	"os"
)

/*
[DEBUG][fwupgrade.cfg: Offset:65536 Size:639]
[DEBUG][kernel: Offset:66175 Size:1082534]
[DEBUG][rootfs: Offset:1148709 Size:2097156]
[DEBUG][fwupgrade.cfg.sig: Offset:3245865 Size:528]
 */

var files = []string{"fwupgrade.cfg", "kernel", "rootfs", "fwupgrade.cfg.sig"}

func main() {
	c, err := tftp.NewClient("192.168.100.8:69")
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, file := range files {
		wt, err := c.Receive(file, "octet")
		if err != nil {
			fmt.Println(err)
		}
		// Optionally obtain transfer size before actual data.
		if n, ok := wt.(tftp.IncomingTransfer).Size(); ok {
			fmt.Printf("Transfer size: %d\n", n)
		}
		n, err := wt.WriteTo(os.Stdout)
		fmt.Printf("%d bytes received\n", n)
		fmt.Printf("%d/%d files received\n", i+1, len(files))
	}
}
