package services

import (
	"Loginmodule/repositories"
	"Loginmodule/vo"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

//Loginhandler
func Loginservice(object *vo.User) string {
	var err error
	var userList vo.Usersobject
	userList.SetUsers(repositories.FindUserByName(object.Getusername()))
	fmt.Println(userList)
	if len(userList.GetUsers()) > 1 || len(userList.GetUsers()) == 0 {
		err = nil
		return vo.Exist
	} else {
		if (userList.GetUsers())[0].Getuserpassword() == object.Getuserpassword() {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"username": object.Getusername(),
				"password": object.Getuserpassword(),
			})
			tokenString, err := token.SignedString([]byte("secret"))
			if err != nil {
				fmt.Println(err)
			}
			object.Settoken(tokenString)
			repositories.Settokendb(object.Getusername(), tokenString)
			tokenString, error := token.SignedString([]byte("secret"))
			if error != nil {
				fmt.Print(error)
			}
			return tokenString
		} else {
			err = nil
			fmt.Print(err)
			return vo.Mismatch
		}
	}
}
