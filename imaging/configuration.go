package imaging

import (
	"net"
	"io/ioutil"
	"strings"
	"errors"
	"crypto/md5"
	"encoding/hex"
	"firstinspires.org/radioconfigtool/util"
)

// Open a tcp socket, and send the config string to the radio.
func SendConfiguration(data string) error {
	conn, err := net.Dial("tcp", "192.168.1.1:8888")
	defer conn.Close()
	if err != nil {
		return err
	}
	valid := checkRadioResponse(conn)
	if !valid {
		return errors.New("Radio replied with an invalid response.")
	}

	_, err = conn.Write([]byte(data))
	return err
}

func checkRadioResponse(conn net.Conn) bool {
	ret := false
	result, err := ioutil.ReadAll(conn)
	if err != nil {
		panic(err)
	}

	dec := DecryptConfigString(string(result))

	ret = IsValid(dec)
	return ret
}

func IsValid(str string) bool {
	/*
	Router reply:
	Line0: Version of router image
	Line1: Hash of version
	Line2: Image build timestamp
	Line3: Image build timestamp Hash
	Line4: Config timestamp
	Line5: Config timestamp hash
	 */
	lines := strings.Split(str, "\n")

	for i := 0; i <= 4; i += 2 {
		md := md5.New()
		md.Write([]byte(lines[i]))
		check := hex.EncodeToString(md.Sum(nil))
		if check != lines[i+1] {
			util.Debug("Checksum failed on router reply (intentionally?): " + str)
			return false
		}
	}
	return true
}
