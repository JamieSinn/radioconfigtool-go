package imaging

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"

	"firstinspires.org/radioconfigtool/config"
)

var key = config.ENCRYPTION_KEY

func EncryptConfigString(str string) string {

	// Mode, Team Number, SSID, WPAKEY, Bandwidth limit, Firewall, DHCP Mode, 2.4GHz Radio ID, 5GHz Radio ID, Comment,\n(termination)
	plaintext := []byte(str)
	fmt.Printf("%s\n", plaintext)
	ciphertext, err := encrypt(key, plaintext)
	if err != nil {
		log.Fatal(err)
	}

	// Base64 Encode the encrypted string to reduce the length
	sEnc := base64.StdEncoding.EncodeToString([]byte(ciphertext))

	return sEnc
}

func DecryptConfigString(str string) string {

	sDec, _ := base64.StdEncoding.DecodeString(str)
	fmt.Printf("%0x\n", sDec)

	result, err := decrypt(key, sDec)
	if err != nil {
		log.Fatal(err)
	}

	return string(result)
}

func encrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return ciphertext, nil
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
