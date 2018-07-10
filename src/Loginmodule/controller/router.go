package controller

import (
	"github.com/gorilla/mux"
)

func Createrouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/login", Loginhandler).Methods("POST")
	router.HandleFunc("/signup", Signuphandler).Methods("POST")
	router.HandleFunc("/logout", Logouthandler).Methods("POST")
	return router
}
