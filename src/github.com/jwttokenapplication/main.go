package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var (
	privateKey []byte
	publicKey  []byte
)

func init() {
	privateKey, _ = ioutil.ReadFile("rsafile.rsa")
	publicKey, _ = ioutil.ReadFile("rsafile.rsa.pub")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	token := jwt.New(jwt.GetSigningMethod("RS256"))
	tokenstring, _ := token.SignedString(privateKey)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, tokenstring)
}

func AuthMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	token, err := jwt.Parse(r.Header.Get("jwt"), func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	if err == nil && token.Valid {
		next(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Println("Go home")
	}
}

func APIHandler(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.WriteHeader(http.StatusOK)
	fmt.Println("WIN")
}

func main() {
	router := mux.NewRouter()
	n := negroni.Classic()
	router.HandleFunc("/login", loginHandler)
	router.Handle("/api", negroni.New(negroni.HandlerFunc(AuthMiddleware), negroni.HandlerFunc(APIHandler)))
	n.UseHandler(router)
	http.ListenAndServe(":3000", n)
}
