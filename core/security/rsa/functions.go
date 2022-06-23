package rsa

import (
	"crypto/rand"
	"crypto/rsa"
)

func Encrypt(pub *rsa.PublicKey, rawData []byte) ([]byte, error) {
	rng := rand.Reader
	return rsa.EncryptPKCS1v15(rng, pub, rawData)
}

func Decrypt(priv *rsa.PrivateKey, encryptedData []byte) ([]byte, error) {
	rng := rand.Reader
	return rsa.DecryptPKCS1v15(rng, priv, encryptedData)
}
