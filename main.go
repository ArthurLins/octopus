package main

import (
	"net/http"

	"github.com/ArthurLins/octopus/core/logger"
	"github.com/ArthurLins/octopus/core/security/rsa"
)

func main() {
	// http.HandleFunc("/", getRoot)

	// err := http.ListenAndServe(":3333", nil)

	// if err != nil {
	// 	log.Fatalln(err)
	// 	return
	// }

	pub, priv, _ := rsa.GenerateKeyPair()

	logger.Info(string(rsa.EncodePublicKey(pub)))
	logger.Info(string(rsa.EncodePrivateKey(priv)))
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()

	//w.
	//fmt.Printf("%s: got / request\n", ctx.Value(keyServerAddr))
	// io.WriteString(w, "AAAAAAAAAA")
}
