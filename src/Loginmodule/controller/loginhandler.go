package controller

import (
	"Loginmodule/services"
	"Loginmodule/vo"
	"encoding/json"
	"log"
	"net/http"
)

//Loginhandler
func Loginhandler(w http.ResponseWriter, r *http.Request) {
	header := r.Header
	log.Println("Header: ")
	log.Println(header)
	object := vo.CreateUserObj(header["Username"][0], header["Password"][0])
	log.Println("Object created: ")
	log.Println(object)
	tokenString := services.Loginservice(&object)
	if tokenString != vo.Null && tokenString != vo.Mismatch && tokenString != vo.Exist {
		json.NewEncoder(w).Encode(tokenString)
		log.Println("Token created: ")
		log.Println(tokenString)
	} else {
		log.Println("Token created: ")
		log.Println(tokenString)
	}
}
