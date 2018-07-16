package controller

import (
	"Loginmodule/RowMappers"
	"Loginmodule/repositories"
	"Loginmodule/services"
	"Loginmodule/vo"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//Signuphandler
func Signuphandler(w http.ResponseWriter, r *http.Request) {
	var user vo.User
	dbuser := RowMappers.MapusertoDbuser(&user)
	err := json.NewDecoder(r.Body).Decode(&dbuser)
	if err != nil {
		log.Panic(err)
	}
	token := services.Signupservice(&dbuser)
	if token != vo.Null && token != vo.Exist {
		json.NewEncoder(w).Encode(token)
		repositories.Settokendb(dbuser.Username, token)
	} else {
		log.Println(token)
		fmt.Println(token)
	}
}
