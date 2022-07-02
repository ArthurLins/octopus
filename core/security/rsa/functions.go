package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

func Encrypt(pub *rsa.PublicKey, rawData []byte) ([]byte, error) {
	rng := rand.Reader
	return rsa.EncryptOAEP(sha256.New(), rng, pub, rawData, nil)
}

func Decrypt(priv *rsa.PrivateKey, encryptedData []byte) ([]byte, error) {
	rng := rand.Reader
	return rsa.DecryptOAEP(sha256.New(), rng, priv, encryptedData, nil)
}
