package fileio

import (
	"os"
	"firstinspires.org/radioconfigtool/util"
	"encoding/csv"
	"bufio"
	"io"
	"fmt"
)

var (
	wpaKeys = make(map[string]string)
)

func LoadKeys() {
	f, err := os.Open("teams.csv")
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

func GetTeamKey(team string) string {
	return wpaKeys[team]
}
