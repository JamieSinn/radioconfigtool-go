package imaging

import (
	"firstinspires.org/radioconfigtool/fileio"
	"strconv"
	"firstinspires.org/radioconfigtool/util"
	"firstinspires.org/radioconfigtool/resources"
	"io"
	"bytes"
	"github.com/pin/tftp"
	"firstinspires.org/radioconfigtool/netconfig"
)

// RouterConfiguration is used for building the configuration string that is sent to the router.
type RouterConfiguration struct {
	Mode        string
	Team        string
	SSID        string
	WPAKey      string
	Firewall    bool
	DHCPEnabled bool
	BWLimit     bool
	RadioID_24  int
	RadioID_5   int
	Event       string
}

// Mode, Team Number, SSID, WPAKEY, Bandwidth limit, Firewall, DHCP Mode, 2.4GHz Radio ID, 5GHz Radio ID, Event,\n(termination)
func (conf RouterConfiguration) BuildConfigString() string {
	return conf.Mode + "," + conf.Team + "," + conf.SSID + "," +
		conf.WPAKey + "," + util.BoolToStr(conf.BWLimit) + "," + util.BoolToStr(conf.Firewall) + "," +
		util.BoolToStr(conf.DHCPEnabled) + "," + strconv.Itoa(conf.RadioID_24) + "," + strconv.Itoa(conf.RadioID_5) + "," +
		conf.Event + ",\n"
}

// RobotRouter is used for the imaging/flashing procedure to hold the information regarding the image file.
type RobotRouter struct {
	Image     fileio.RouterImage
	Model     string
	ARPString string
	ConfigIP  []byte
}

func (router RobotRouter) ReadHandler(filename string, rf io.ReaderFrom) error {

	file, err := router.Image.GetFile(filename)
	if err != nil {
		util.Debug("Could not find requested file: " + filename)
		return err
	}
	filecount := <-netconfig.SentFiles
	filecount++
	netconfig.SentFiles <- filecount
	rf.(tftp.OutgoingTransfer).SetSize(int64(file.Size))
	n, err := rf.ReadFrom(bytes.NewReader(file.Data))
	if err != nil {
		util.Debug("%v", err)
		return err
	}
	util.Debug("%d bytes sent", n)
	return nil
}

func (router RobotRouter) VerifyImage() {
	data, err := resources.Asset(router.Image.Path)

	if err != nil {
		panic("Could not find image for " + router.Model + " at " + router.Image.Path)
	}

	if fileio.VerifyImage(data, router.Image, router.Image.ExpectedSize) {
		panic("Could not verify image for " + router.Model + " at " + router.Image.Path)
	}
}

// Type for array of all available radios
type RobotRadioList []RobotRouter