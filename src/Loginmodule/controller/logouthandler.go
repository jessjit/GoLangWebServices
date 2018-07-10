package controller

import (
	"Loginmodule/repositories"
	"Loginmodule/services"
	"Loginmodule/vo"
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/mitchellh/mapstructure"
)

func Logouthandler(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	valid := services.Logoutservice(auth)
	fmt.Println("Valid")
	fmt.Println(valid)
	context.Set(r, "decoded", valid.Claims)
	decoded := context.Get(r, "decoded")
	fmt.Println("decoded")
	fmt.Println(decoded)
	var userr vo.DbUser
	mapstructure.Decode(decoded.(jwt.MapClaims), &userr)
	tokenString, error := valid.SignedString([]byte("secret"))
	userr.Token = tokenString
	if error != nil {
		panic(error)
	}
	fmt.Println(userr)
	fmt.Println(userr.Username)
	var userList vo.Usersobject
	userList.SetUsers(repositories.FindUserByName(userr.Username))
	fmt.Println("Userlist")
	fmt.Println(userList)
	if len(userList.GetUsers()) != 1 {
		fmt.Println("error")
	} else {
		repositories.Removedb(userList.GetUsers()[0].Getuserid())
		fmt.Println("logout success")
	}
}
