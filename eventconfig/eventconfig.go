package eventconfig

import (
	"os"
	"fmt"
	"net/http"
	"io"
	"firstinspires.org/radioconfigtool/netconfig"
	"strings"
	"bytes"
	"firstinspires.org/radioconfigtool/config"
)

func getImage(fileName string) {
	netconfig.ResetNetworkAdapter(config.NETINT_WLAN)

	url := config.FMSUrl + fileName

	output, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error while creating", fileName, "-", err)
		return
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}
	defer response.Body.Close()

	_, err = io.Copy(output, response.Body)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}
}

func GetTeams() []string {
	url := config.FMSUrl + "teams.txt"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return []string{}
	}
	defer response.Body.Close()

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(response.Body)
	teams := strings.Split(buffer.String(), "\n")
	return teams
}

func GetAllImages(teams []string) {

	for _, team := range teams {
		getImage(team + "-AN.bin")
		getImage(team + "-AC.bin")
	}
}
