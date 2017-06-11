package imaging

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"log"

	"firstinspires.org/radioconfigtool/config"
)

var key = []byte(config.ENCRYPTION_KEY)

// EncryptConfigString takes a raw config string, encrypts it using AES-256, then base64
// encodes that to make the string size smaller
func EncryptConfigString(str string) string {

	plaintext := []byte(str)
	ciphertext, err := encrypt(key, plaintext)
	if err != nil {
		log.Fatal(err)
	}

	// Base64 Encode the encrypted string to reduce the length
	sEnc := base64.StdEncoding.EncodeToString([]byte(ciphertext))

	return sEnc
}

// DecryptConfigString decodes from a base64 string, and then decrypts the AES-256 string
// and returns the raw config string.
func DecryptConfigString(str string) string {

	sDec, _ := base64.StdEncoding.DecodeString(str)

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
