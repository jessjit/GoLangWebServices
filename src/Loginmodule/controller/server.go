package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Runserver(r *mux.Router) {
	log.Println("Server running")
	log.Fatal(http.ListenAndServe(":8000", r))
}
