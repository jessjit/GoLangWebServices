package services

import (
	"Loginmodule/repositories"
	"Loginmodule/vo"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

//Loginhandler
func Loginservice(object *vo.User) string {
	fmt.Print(object)
	fmt.Println(object.Getusername())
	var err error
	var userList vo.Usersobject
	userList.SetUsers(repositories.FindUserByName(object.Getusername()))
	fmt.Println(len("userList.GetUsers()"))
	fmt.Println(userList.GetUsers())
	if len(userList.GetUsers()) > 1 || len(userList.GetUsers()) == 0 {
		err = nil
		return "no such user exists"
		//panic(err)
	} else {
		fmt.Print("userList.GetUsers()[0].Getuserpassword()")
		fmt.Println(userList.GetUsers())

		fmt.Print("object.Getuserpassword()")
		fmt.Println(object.Getuserpassword())
		if (userList.GetUsers())[0].Getuserpassword() == object.Getuserpassword() {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"username": object.Getusername(),
				"password": object.Getuserpassword(),
			})
			object.Settoken(token)
			tokenString, error := token.SignedString([]byte("secret"))
			if error != nil {
				fmt.Print(error)
			}
			return tokenString
		} else {
			err = nil
			fmt.Print(err)
			return "password mismatch"
			//panic(err)
		}
	}
}
