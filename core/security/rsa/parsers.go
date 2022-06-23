package rsa

import (
	"crypto/rsa"
	"crypto/x509"
)

func DecodePublicKey(encodedPublicKey []byte) {
	x509.ParsePKCS1PublicKey(encodedPublicKey)
}

func DecodePrivateKey(encodedPrivateKey []byte) {
	x509.ParsePKCS1PrivateKey(encodedPrivateKey)
}

func EncodePublicKey(pub *rsa.PublicKey) []byte {
	return x509.MarshalPKCS1PublicKey(pub)
}

func EncodePrivateKey(priv *rsa.PrivateKey) []byte {
	return x509.MarshalPKCS1PrivateKey(priv)
}
