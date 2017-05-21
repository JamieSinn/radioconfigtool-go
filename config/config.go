package config

var (
	// EventMode is toggled when at competition
	EventMode = false
	// PracticeFieldMode is only used for events with official practice fields (Championships, District Championships)
	PracticeFieldMode = false

)

const (
	//FMSUrl is the url to download team images when in eventmode.
	FMSUrl = "http://10.0.100.5/imaging/images/"

	// NETINT_LAN is the interface name to use for imaging/configuration
	NETINT_LAN = "LAN"

	//NETINT_WLAN is the interface name that is used to connect to the FMS
	NETINT_WLAN = "WLAN"

	// DEBUG - Whether or not to print debug messages
	DEBUG = true
)