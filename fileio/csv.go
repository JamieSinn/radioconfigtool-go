package fileio

import (
	"os"
	"firstinspires.org/radioconfigtool/util"
	"encoding/csv"
	"bufio"
	"io"
	"firstinspires.org/radioconfigtool/config"
)

var (
	wpaKeys = make(map[string]string)
)

// LoadKeys loads the WPA Keys that the FMS exports during the Event Manager setup.
func LoadKeys() {
	f, err := os.Open(config.WPA_KEY_PATH)
	if err != nil {
		util.Debug(err)
		return
	}

	r := csv.NewReader(bufio.NewReader(f))

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}
		util.Debug(record)
		wpaKeys[record[0]] = record[1]
	}
}

// GetTeamKey is an exported access to the wpaKeys map
func GetTeamKey(team string) string {
	return wpaKeys[team]
}
