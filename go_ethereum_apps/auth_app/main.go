package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/rs/cors"
	"github.com/spruceid/siwe-go"
	"net/http"
)

type signInParams struct {
	Message   string `json:"message"`
	Signature string `json:"signature"`
}

var store = sessions.NewCookieStore([]byte("SESSION_KEY"))

func main() {
	mux := http.NewServeMux()

	store.Options.HttpOnly = true

	mux.HandleFunc("GET /api/nonce", makeHTTPHandleFunc(getNonce))
	mux.HandleFunc("POST /api/signin", makeHTTPHandleFunc(signin))
	mux.HandleFunc("GET /api/participate", withAuth(makeHTTPHandleFunc(participate)))
	mux.HandleFunc("GET /api/signout", withAuth(makeHTTPHandleFunc(signout)))

	//handler := cors.Default().Handler(mux)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		//Debug:            true,
	})

	handler := c.Handler(mux)

	fmt.Println("server started on 8085")
	http.ListenAndServe(":8085", handler)
}

func participate(w http.ResponseWriter, r *http.Request) error {

	// todo: get data from store

	return WriteText(w, http.StatusOK, "participate")
}

func getNonce(w http.ResponseWriter, r *http.Request) error {
	nonce := siwe.GenerateNonce()

	session, _ := store.Get(r, "sessionId")
	session.Values["nonce"] = nonce

	if err := session.Save(r, w); err != nil {
		return err
	}

	return WriteText(w, http.StatusOK, nonce)
}

func signin(w http.ResponseWriter, r *http.Request) error {
	var data signInParams
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return err
	}

	siweMessage, err := siwe.ParseMessage(data.Message)
	if err != nil {
		return err
	}

	session, _ := store.Get(r, "sessionId")
	nonce := session.Values["nonce"]

	if siweMessage.GetNonce() != nonce {
		return errors.New("message nonce doesn't match")
	}

	if _, err := siweMessage.VerifyEIP191(data.Signature); err != nil {
		return err
	}

	session.Values["address"] = siweMessage.GetAddress().String()
	session.Values["authenticated"] = true

	if err := session.Save(r, w); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, nil)

}

func signout(w http.ResponseWriter, r *http.Request) error {
	session, _ := store.Get(r, "sessionId")
	session.Values["authenticated"] = false
	delete(session.Values, "address")

	if err := session.Save(r, w); err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, nil)
}

func withAuth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "sessionId")

		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			WriteJSON(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		handlerFunc(w, r)
	}
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, err.Error())
		}
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func WriteText(w http.ResponseWriter, status int, v string) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(v))

	return nil
}
