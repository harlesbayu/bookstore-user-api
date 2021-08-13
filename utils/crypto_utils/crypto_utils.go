package crypto_utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"io"
)

var secretKey = []byte("zfFmdquVqYnsK3gVnanAqtZANIrwWN1l")

func hashPassword(p []byte) []byte {
	var salt = []byte("secret-1234")

	h := sha256.New()
	h.Write(append(p, salt...))
	return h.Sum(nil)
}

func encodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func decodeBase64(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

func EncryptPassword(password string) string {
	block, err := aes.NewCipher(hashPassword(secretKey))
	if err != nil {
		panic(err)
	}
	ciphertext := make([]byte, aes.BlockSize+len([]byte(password)))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(password))
	return encodeBase64(ciphertext)
}

func DecryptPassword(password string) string {
	text := decodeBase64(password)
	block, err := aes.NewCipher(hashPassword(secretKey))
	if err != nil {
		panic(err)
	}
	if len(text) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	return string(text)
}
