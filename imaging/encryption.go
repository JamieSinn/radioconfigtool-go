package imaging

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"firstinspires.org/radioconfigtool/config"
)

var key = []byte(config.ENCRYPTION_KEY)

// EncryptConfigString takes a raw config string, encrypts it using AES-256 CFB, then base64
// encodes that to make the string size smaller
func EncryptConfigString(str string) string {

	plaintext := []byte(str)
	ciphertext, err := encrypt(key, plaintext)
	if err != nil {
		panic(err)
	}
	// Base64 Encode the encrypted string to reduce the length
	sEnc := base64.StdEncoding.EncodeToString([]byte(ciphertext))
	return sEnc
}

// EncryptConfigStringCBC takes a raw config string, encrypts it using AES-256 CBC, then base64
// encodes that to make the string size smaller
func EncryptConfigStringCBC(str string) string {

	// CBC Needs padding to a multiple of the block size (16)
	if len(str)%16 != 0 {
		// Round to nearest multiple, then subtract the current length to get the final amount to add.
		count := ((len(str) + 15) / 16 * 16) - len(str)
		for i := 0; i < count; i++ {
			str += "="
		}
	}
	plaintext := []byte(str)
	ciphertext := encryptCBC(key, plaintext)
	// Base64 Encode the encrypted string to reduce the length
	sEnc := base64.StdEncoding.EncodeToString([]byte(ciphertext))
	return sEnc
}

// DecryptConfigString decodes from a base64 string, and then decrypts the AES-256 CFB string
// and returns the raw config string.
func DecryptConfigString(str string) string {

	sDec, _ := base64.StdEncoding.DecodeString(str)

	result, err := decrypt(key, sDec)
	if err != nil {
		panic(err)
	}

	return string(result)
}

// DecryptConfigString decodes from a base64 string, and then decrypts the AES-256 CFB string
// and returns the raw config string.
func DecryptConfigStringCBC(str string) string {

	sDec, _ := base64.StdEncoding.DecodeString(str)

	result := decryptCBC(key, sDec)

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

func decryptCBC(key, ciphertext []byte) string {
	//ciphertext, _ = hex.DecodeString("f363f3ccdcb12bb883abf484ba77d9cd7d32b5baecb3d4b1b3e0e4beffdb3ded")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	// CBC mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}
	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(ciphertext, ciphertext)

	// If the original plaintext lengths are not a multiple of the block
	// size, padding would have to be added when encrypting, which would be
	// removed at this point. For an example, see
	// https://tools.ietf.org/html/rfc5246#section-6.2.3.2. However, it's
	// critical to note that ciphertexts must be authenticated (i.e. by
	// using crypto/hmac) before being decrypted in order to avoid creating
	// a padding oracle.

	return string(ciphertext)
}

func encryptCBC(key, plaintext []byte) string {

	// CBC mode works on blocks so plaintexts may need to be padded to the
	// next whole block. For an example of such padding, see
	// https://tools.ietf.org/html/rfc5246#section-6.2.3.2. Here we'll
	// assume that the plaintext is already of the correct length.
	if len(plaintext)%aes.BlockSize != 0 {
		panic("plaintext is not a multiple of the block size")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.
	return string(ciphertext)
}
