package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"encoding/base64"
	"strings"
	"crypto/aes"
	"errors"
	"crypto/cipher"
	"log"
)

const (
	competition = "dsS`p%%Tu\"zT70=F*Mm7ivx[T~Mr@HcX"
	home = "IzLNm4rZK77TBCXopuRhufEP7x6UBOWl"
)

func main() {

	fmt.Println("Creating FRC Radio...")
	ln, err := net.Listen("tcp", "192.168.1.1:8888")
	if err != nil {
		panic(err)
	}
	fmt.Println("Handling incoming connections...")
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	if conn == nil {
		return
	}
	fmt.Println("Connection from: " + conn.RemoteAddr().String())
	outgoing := "Model: 1\nVersion: 2017.2\nEvent: \n"
	enc := base64.StdEncoding.EncodeToString([]byte(outgoing))
	fmt.Println("Data:\n " + outgoing)
	fmt.Println("Encoded: " + enc)

	conn.Write([]byte(enc))

	fmt.Println("Waiting for connection data....")
	result, err := ioutil.ReadAll(conn)
	if err != nil {
		panic(err)
	}
	raw := string(result)
	split := strings.Split(raw, ",")
	key := ""
	switch split[0] {
	case "0":
		fmt.Println("Not at competition, use the home key for decryption")
		key = home
		break
	case "1":
		fmt.Println("At competition, using event key for decryption")
		key = competition
		break
	}
	b64 := split[1]

	fmt.Println("Encoded + Encrypted: " + b64)
	dec, _ := base64.StdEncoding.DecodeString(b64)
	fmt.Println("Encrypted: " + string(dec))

	dEncr, err := decrypt([]byte(key), dec)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Decrypted: " + string(dEncr))
}

func decrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(text) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil, err
	}
	return data, nil
}

