package services

import (
	"Loginmodule/repositories"
	"Loginmodule/vo"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

func Signupservice(user *vo.User) string {
	var userList vo.Usersobject
	userList.SetUsers(repositories.FindUserByName(user.Getusername()))
	if len(userList.GetUsers()) > 0 {
		return "user already exits"
	}
	b := repositories.Insertdb(user)
	if b == true {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": user.Getusername(),
			"password": user.Getuserpassword(),
		})
		user.Settoken(token)
		tokenString, error := token.SignedString([]byte("secret"))
		if error != nil {
			fmt.Println(error)
		}
		return tokenString
	} else {
		return "null"
	}
}
