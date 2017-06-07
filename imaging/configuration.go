package imaging

import (
	"net"
	"io/ioutil"
	"strings"
)

func SendConfiguration(data string) {
	conn, err := net.Dial("tcp", "192.168.1.1:8888")
	defer conn.Close()
	if err != nil {
		// Could not connect
	}
	valid := checkRadioResponse(conn)
	if !valid {
		//Invalid radio response
	}
	//Send config string

	conn.Write([]byte(data))
}

// TODO: Change based on radio config rewrite
func checkRadioResponse(conn net.Conn) bool {
	ret := false
	result, err := ioutil.ReadAll(conn)
	if err != nil {
		panic(err)
	}
	split := strings.Split(string(result), "\n")

	if strings.HasPrefix(split[0], "Team:") &&
		strings.HasPrefix(split[1], "Event:") {
		ret = true
	}
	return ret
}
