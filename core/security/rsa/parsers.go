package rsa

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

func DecodePublicKey(encodedPublicKey []byte) {
	x509.ParsePKCS1PublicKey(encodedPublicKey)
	//Todo::
}

func DecodePrivateKey(encodedPrivateKey []byte) {
	x509.ParsePKCS1PrivateKey(encodedPrivateKey)
	//Todo::
}

func EncodePublicKey(pub *rsa.PublicKey) []byte {
	encoded := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(pub),
	})
	return encoded
}

func EncodePrivateKey(priv *rsa.PrivateKey) []byte {
	encoded := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(priv),
	})
	return encoded
}
