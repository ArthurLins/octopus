package main

import (
	"net/http"

	"github.com/ArthurLins/octopus/core/logger"
	"github.com/ArthurLins/octopus/core/security/aes"
)

func main() {
	// http.HandleFunc("/", getRoot)

	// err := http.ListenAndServe(":3333", nil)

	// if err != nil {
	// 	log.Fatalln(err)
	// 	return
	// }

	dt, err := aes.EncryptBase64([]byte("key"), []byte("Olá"))

	if err != nil {
		logger.Error(err.Error())
	}

	logger.Error(dt)

	dt2, err := aes.DecryptBase64([]byte("key"), dt)

	if err != nil {
		logger.Error("AA: " + err.Error())

	}

	logger.Info(string(dt))
	logger.Info(string(dt2))
	//aes.Encrypt([]byte("aaaa"), []byte("Olá"))

	//pub, priv, _ := rsa.GenerateKeyPair()

	//logger.Info(string(rsa.EncodePublicKey(pub)))
	//logger.Info(string(rsa.EncodePrivateKey(priv)))
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()

	//w.
	//fmt.Printf("%s: got / request\n", ctx.Value(keyServerAddr))
	// io.WriteString(w, "AAAAAAAAAA")
}
