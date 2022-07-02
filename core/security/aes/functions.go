package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"io"
)

//KEY = SHA256(INPUT-KEY)

func Encrypt(key []byte, data []byte) ([]byte, error) {

	keySum := sha256.Sum256(key)
	slice := keySum[:]

	c, err := aes.NewCipher(slice)

	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)

	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, data, nil), nil
}

func EncryptBase64(key []byte, data []byte) (string, error) {
	data, err := Encrypt(key, data)
	return base64.StdEncoding.EncodeToString(data), err
}

func Decrypt(key []byte, encryptedData []byte) ([]byte, error) {

	keySum := sha256.Sum256(key)
	slice := keySum[:]

	c, err := aes.NewCipher(slice)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(encryptedData) < nonceSize {
		return nil, err
	}

	nonce, cData := encryptedData[:nonceSize], encryptedData[nonceSize:]

	data, err := gcm.Open(nil, nonce, cData, nil)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func DecryptBase64(key []byte, encryptedBase64 string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(encryptedBase64)
	if err != nil {
		return nil, err
	}
	return Decrypt(key, data)
}
