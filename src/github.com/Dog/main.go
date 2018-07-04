package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Dog/Service"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	fmt.Println("Starting the application")
	router.HandleFunc("/intro", Service.Introhandler)
	router.HandleFunc("/login", Service.Loginhandler).Methods("POST")
	router.HandleFunc("/signup", Service.Signuphandler).Methods("POST")
	router.HandleFunc("/logout", Service.Logouthandler)
	log.Fatal(http.ListenAndServe(":8008", router))
}
