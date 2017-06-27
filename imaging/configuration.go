package imaging

import (
	"net"
	"strings"
	"errors"
	"firstinspires.org/radioconfigtool/util"
	"encoding/base64"
	"firstinspires.org/radioconfigtool/config"
	"strconv"
	"time"
	"io"
)

// Open a tcp socket, and send the config string to the radio.
func SendConfiguration(data string) error {
	conn, err := net.Dial("tcp", "192.168.1.1:8888")
	if err != nil {
		util.Debug(err)
		return err
	}
	buf := make([]byte, 0, 4096)
	tmp := make([]byte, 256)
	for {
		n, err := conn.Read(tmp)
		if err != nil {
			if err != io.EOF {
				util.Debug("Socket: read error:", err)
			}
			break
		}
		util.Debug("Socket: got", n, "bytes.")
		util.Debug("'" + string(tmp) + "'")
		buf = append(buf, tmp[:n]...)
		if n < len(tmp) {
			break
		}
	}
	result := buf
	util.Debug("Raw data from connection: '" + string(result) + "'")
	// String is a base64 encoded string to prevent immediate prying eyes.
	// Must decode before use.
	stripped := strings.Replace(string(result), "\n", "", -1)
	dec, _ := base64.StdEncoding.DecodeString(stripped)

	util.Debug("Decoded string: " + string(dec))

	// Split all lines for easier use
	lines := strings.Split(string(dec), "\n")
	for _, l := range lines {
		util.Debug("Line: '" + l + "'")
	}
	// Check if the data received is valid
	valid := IsValid(lines)

	if !valid {
		return errors.New("Invalid")
	}

	util.Debug("Data is valid.")
	// Check if the radio image version is up to date.
	// The first check is for debugging only.
	uptodate := (config.DEBUG && lines[1] == "Version: ") || IsUpToDate(lines[1])

	if !uptodate {
		return errors.New("OutOfDate")
	}

	util.Debug("Version is up to date.")
	// Check if the radio is attempted to be programmed while an event is going on
	// by the team configuration build.
	atevent := config.EventMode() && IsWithinCompetition(lines[2])

	if atevent {
		return errors.New("AtEvent")
	}
	util.Debug("Not at an event.")

	// Prepend the configuration string with a 1|0 if this is for competition.
	// This is because the radio programmer needs to use the correct key to decrypt the string.
	data = util.BoolToStr(config.EventMode()) + "," + data + ",\n"
	util.Debug("To Send: '" + data + "'")
	_, err = conn.Write([]byte(data))
	return err
}

func IsValid(lines []string) bool {
	/*
	Router reply:
	Line0: Model: {Hardware Enum}
	Line1: Version: {OpenWRT Build Version}
	Line2: Event: {EventInfo}
	 */

	if strings.HasPrefix(lines[0], "Model: ") &&
		strings.HasPrefix(lines[1], "Version: ") &&
		strings.HasPrefix(lines[2], "Event: ") {
		return true
	}
	return false
}

func IsUpToDate(verstr string) bool {

	raw := verstr[9:]

	major, err := strconv.Atoi(raw[:4])
	if err != nil {
		return false
	}
	minor, err := strconv.Atoi(raw[5:])
	if err != nil {
		return false
	}
	if config.MIN_BUILD_maj <= major &&
		config.MIN_BUILD_min <= minor {
		return true
	}
	return false
}

func IsWithinCompetition(event string) bool {
	if len(event) == 7 {
		return false
	}
	raw := event[7:]

	epoch, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		return false
	}

	tm := time.Unix(epoch, 0)
	if tm.Unix() > time.Now().Unix() {
		return true
	}
	return false
}
