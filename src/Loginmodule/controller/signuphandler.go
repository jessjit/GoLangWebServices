package controller

import (
	"Loginmodule/services"
	"Loginmodule/vo"
	"encoding/json"
	"fmt"
	"net/http"
)

func Signuphandler(w http.ResponseWriter, r *http.Request) {
	var user vo.User
	err := json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)
	if err != nil {
		panic(err)
	}
	token := services.Signupservice(&user)
	if token != "null" && token != "user already exits" {
		json.NewEncoder(w).Encode(token)
	} else {
		fmt.Println(token)
	}
}
