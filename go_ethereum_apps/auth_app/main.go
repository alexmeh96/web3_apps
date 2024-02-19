package main

import (
	"encoding/json"
	"fmt"
	"github.com/spruceid/siwe-go"
	"net/http"
)

type signInParams struct {
	Message   string `json:"message"`
	Signature string `json:"signature"`
}

var sessionStore = make(map[string]string)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/nonce/{address}", func(w http.ResponseWriter, r *http.Request) {
		address := r.PathValue("address")
		fmt.Printf("address: %s", address)

		nonce := siwe.GenerateNonce()
		sessionStore[address] = nonce

		w.Header().Set("Access-Control-Allow-Origin", "*")

		w.Write([]byte(nonce))
	})

	mux.HandleFunc("OPTIONS /*", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Max-Age", "300")
	})

	mux.HandleFunc("POST /api/verify", func(w http.ResponseWriter, r *http.Request) {
		var data signInParams
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		siweMessage, err := siwe.ParseMessage(data.Message)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		address := siweMessage.GetAddress()

		nonce := sessionStore[address.String()]
		// todo: вытащить nonce из ссесии по адрессу
		//nonce := siweMessage.GetNonce()

		if siweMessage.GetNonce() != nonce {
			http.Error(w, "Message nonce doesn't match", http.StatusBadRequest)
			return
		}

		publicKey, err := siweMessage.VerifyEIP191(data.Signature)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println(publicKey)

		w.Header().Set("Access-Control-Allow-Origin", "*")
	})

	fmt.Println("server started on 8086")
	http.ListenAndServe(":8085", mux)
}
