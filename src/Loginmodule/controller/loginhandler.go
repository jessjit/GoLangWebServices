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
	object := vo.CreateUserObj(header["Username"][0], header["Password"][0])
	fmt.Println(object)
	tokenString := services.Loginservice(&object)
	if tokenString != vo.Null && tokenString != vo.Mismatch && tokenString != vo.Exist {
		json.NewEncoder(w).Encode(tokenString)
	} else {
		fmt.Println(tokenString)
	}
}
