package config

var (
	// EventMode is toggled via go build. Used to enable or disable certain features for competition.
	eventmode = "false"

	// Encryption key for sending to radio.
	ENCRYPTION_KEY = "dsS`p%%Tu\"zT70=F*Mm7ivx[T~Mr@HcX"
)

const (
	// NETINT_LAN is the interface name to use for imaging/configuration
	NETINT_LAN = "LAN"

	// DEBUG - Whether or not to print debug messages
	DEBUG = true

	// WPA_KEY_PATH - Where to import keys from during event.
	WPA_KEY_PATH = "keys.csv"
)

func EventMode() bool {
	return eventmode == "true"
}
