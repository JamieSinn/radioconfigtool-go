package imaging

import (
	"net"
	"io/ioutil"
	"strings"
	"errors"
	"firstinspires.org/radioconfigtool/util"
	"encoding/base64"
	"firstinspires.org/radioconfigtool/config"
	"strconv"
	"time"
)

// Open a tcp socket, and send the config string to the radio.
func SendConfiguration(data string) error {
	conn, err := net.Dial("tcp", "192.168.1.1:8888")
	defer conn.Close()
	if err != nil {
		return err
	}

	result, err := ioutil.ReadAll(conn)
	if err != nil {
		panic(err)
	}
	// String is a base64 encoded string to prevent immediate prying eyes.
	// Must decode before use.
	dec, _ := base64.StdEncoding.DecodeString(string(result))

	// Split all lines for easier use
	lines := strings.Split(string(dec), "\n")
	// Check if the data received is valid
	valid := IsValid(lines)

	if !valid {
		return errors.New("Invalid")
	}

	// Check if the radio image version is up to date.
	uptodate := IsUpToDate(lines[1])

	if !uptodate {
		return errors.New("OutOfDate")
	}

	// Check if the radio is attempted to be programmed while an event is going on
	// by the team configuration build.
	atevent := IsWithinCompetition(lines[2]) && config.EventMode()

	if atevent {
		return errors.New("AtEvent")
	}


	// Prepend the configuration string with a 1|0 if this is for competition.
	// This is because the radio programmer needs to use the correct key to decrypt the string.
	data = util.BoolToStr(config.EventMode()) + "," + data

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