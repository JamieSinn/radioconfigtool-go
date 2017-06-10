package imaging

import (
	"firstinspires.org/radioconfigtool/fileio"
	"strconv"
	"firstinspires.org/radioconfigtool/util"
	"firstinspires.org/radioconfigtool/resources"
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
	Comment     string
}

// Mode, Team Number, SSID, WPAKEY, Bandwidth limit, Firewall, DHCP Mode, 2.4GHz Radio ID, 5GHz Radio ID, Comment,\n(termination)
func (conf RouterConfiguration) BuildConfigString() string {
	return conf.Mode + "," + conf.Team + "," + conf.SSID + "," +
		conf.WPAKey + "," + util.BoolToStr(conf.BWLimit) + "," + util.BoolToStr(conf.Firewall) + "," +
		util.BoolToStr(conf.DHCPEnabled) + "," + strconv.Itoa(conf.RadioID_24) + "," + strconv.Itoa(conf.RadioID_5) + "," +
		conf.Comment + ",\n"
}

// RobotRouter is used for the imaging/flashing procedure to hold the information regarding the image file.
type RobotRouter struct {
	Image     fileio.RouterImage
	Model     string
	ARPString string
	ConfigIP  []byte
}

// Type for array of all available radios
type RobotRadioList []RobotRouter

func (router RobotRouter) VerifyImage() {
	data, err := resources.Asset(router.Image.Path)

	if err != nil {
		panic("Could not find image for " + router.Model + " at " + router.Image.Path)
	}

	if fileio.VerifyImage(data, router.Image, router.Image.ExpectedSize) {
		panic("Could not verify image for " + router.Model + " at " + router.Image.Path)
	}
}
