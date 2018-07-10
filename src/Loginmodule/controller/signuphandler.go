package controller

import (
	"Loginmodule/RowMappers"
	"Loginmodule/services"
	"Loginmodule/vo"
	"encoding/json"
	"fmt"
	"net/http"
)

func Signuphandler(w http.ResponseWriter, r *http.Request) {
	var user vo.User
	fmt.Println("R's body")
	fmt.Println(r.Body)
	dbuser := RowMappers.MapusertoDbuser(&user)
	err := json.NewDecoder(r.Body).Decode(&dbuser)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
	fmt.Println(dbuser)
	token := services.Signupservice(&dbuser)
	if token != vo.Null && token != vo.Exist {
		json.NewEncoder(w).Encode(token)
	} else {
		fmt.Println(token)
	}
}
