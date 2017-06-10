package config

var (
	// EventMode is toggled via go build. Used to enable or disable certain features for competition.
	eventmode = "false"

	// Encryption key for sending to radio.
	// dsS`p%%Tu"zT70=F*Mm7ivx[T~Mr@HcX
	// Not defined as a string for security
	ENCRYPTION_KEY = []byte{100, 115, 83, 96, 112, 37, 37, 84, 117, 34, 122, 84, 55, 48, 61, 70, 42, 77, 109, 55, 105, 118, 120, 91, 84, 126, 77, 114, 64, 72, 99, 88}
)

const (
	// NETINT_LAN is the interface name to use for imaging/configuration
	NETINT_LAN = "LAN"

	// DEBUG - Whether or not to print debug messages
	DEBUG = true
)

func EventMode() bool {
	return eventmode == "true"
}
