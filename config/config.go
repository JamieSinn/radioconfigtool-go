package config

import "time"

var (
	// EventMode is toggled via go build. Used to enable or disable certain features for competition.
	eventmode = "false"

	// Encryption key for sending to radio. (Event Mode)
	ENCRYPTION_KEY = "dsS`p%%Tu\"zT70=F*Mm7ivx[T~Mr@HcX"
)

const (
	// NETINT_LAN is the interface name to use for imaging/configuration
	NETINT_LAN = "LAN"

	// DEBUG - Whether or not to print debug messages
	DEBUG = true

	// WPA_KEY_PATH - Where to import keys from during event.
	WPA_KEY_PATH = "keys.csv"

	// MIN_BUILD - Minimum build for the OpenWRT image to allow configuring.
	MIN_BUILD_maj = 2017
	MIN_BUILD_min = 2

	// ARP_TIMEOUT - Maximum amount of time to wait for ARP requests.
	ARP_TIMEOUT = 15 * time.Second

	// TFTP_TIMEOUT - Maximum amount of time to send all files
	TFTP_TIMEOUT = 5 * time.Minute
)

func EventMode() bool {
	return eventmode == "true"
}
