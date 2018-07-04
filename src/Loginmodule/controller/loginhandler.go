package controller

import (
	"Loginmodule/services"
	"Loginmodule/vo"
	"encoding/json"
	"fmt"
	"net/http"
)

//Loginhandler

func Loginhandler(w http.ResponseWriter, r *http.Request) {
	header := r.Header
	var object vo.User
	fmt.Print("Header")
	fmt.Println(r.Header)
	object.Setusername(header["Username"][0])
	object.Setpassword(header["Password"][0])
	fmt.Print("object")
	fmt.Println(object)
	token := services.Loginservice(&object)
	if token != "null" && token != "password mismatch" && token != "no such user exists" {
		json.NewEncoder(w).Encode(token)
	} else {
		fmt.Println(token)
	}
}
