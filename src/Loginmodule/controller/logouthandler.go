package controller

import (
	"Loginmodule/repositories"
	"Loginmodule/services"
	"Loginmodule/vo"
	"fmt"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/mitchellh/mapstructure"
)

//Logouthandler
func Logouthandler(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	log.Println("Authorization:")
	log.Println(auth)
	valid := services.Logoutservice(auth)
	context.Set(r, "decoded", valid.Claims)
	decoded := context.Get(r, "decoded")
	var userr vo.DbUser
	mapstructure.Decode(decoded.(jwt.MapClaims), &userr)
	log.Println("User: ")
	log.Println(userr)
	tokenString, error := valid.SignedString([]byte("secret"))
	userr.Token = tokenString
	if error != nil {
		log.Panic(error)
	}
	var userList vo.Usersobject
	userList.SetUsers(repositories.FindUserByName(userr.Username))
	log.Println("Userlist: ")
	log.Println(userList)
	if len(userList.GetUsers()) != 1 {
		log.Panic(error)
		fmt.Println("error")
	} else {
		repositories.Removedb(userList.GetUsers()[0].Getuserid())
		log.Println("Logout successful")
		fmt.Println("Logout successful")
	}
}
