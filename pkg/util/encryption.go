package util

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func EncodeCFB(secret, strToEncode string) (string, error) {
	bytes := []byte(secret)
	block, err := aes.NewCipher(bytes)
	if err != nil {
		return "", err
	}
	plainText := []byte(strToEncode)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func decodeBase64(s string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func DecodeCFB(secret, encoded string) (string, error) {
	bytes := []byte(secret)
	block, err := aes.NewCipher(bytes)
	if err != nil {
		return "", err
	}
	cipherText, err := decodeBase64(encoded)
	if err != nil {
		return "", err
	}
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}
