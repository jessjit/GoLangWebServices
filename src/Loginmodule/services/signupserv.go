package services

import (
	"Loginmodule/repositories"
	"Loginmodule/vo"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

func Signupservice(dbuser *vo.DbUser) string {
	var userList vo.Usersobject
	userList.SetUsers(repositories.FindUserByName(dbuser.Username))
	if len(userList.GetUsers()) > 0 {
		return vo.Exist
	}
	b := repositories.Insertdb(dbuser)
	if b == true {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": dbuser.Username,
			"password": dbuser.Password,
		})
		tokenString, err := token.SignedString([]byte("secret"))
		dbuser.Token = tokenString
		if err != nil {
			panic(err)
		}
		tokenString, error := token.SignedString([]byte("secret"))
		if error != nil {
			fmt.Println(error)
		}
		return tokenString
	} else {
		return vo.Null
	}
}
