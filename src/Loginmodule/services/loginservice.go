package services

import (
	"Loginmodule/repositories"
	"Loginmodule/vo"
	"fmt"
	"strings"

	"log"

	jwt "github.com/dgrijalva/jwt-go"
)

//Loginhandler passes control
func Loginservice(object *vo.User) string {
	var err error
	var userList vo.Usersobject
	userList.SetUsers(repositories.FindUserByName(object.Getusername()))
	if len(userList.GetUsers()) > 1 || len(userList.GetUsers()) == 0 {
		log.Println(vo.Exist)
		return vo.Exist
	} else {
		if (userList.GetUsers())[0].Getuserpassword() == object.Getuserpassword() {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"username": object.Getusername(),
				"password": object.Getuserpassword(),
			})
			tokenString, err := token.SignedString([]byte("secret"))
			if err != nil {
				log.Panic(err)
				fmt.Println(err)
			}
			log.Println(tokenString)
			object.Settoken(tokenString)
			repositories.Settokendb(object.Getusername(), tokenString)
			tokenString, error := token.SignedString([]byte("secret"))
			if error != nil {
				log.Panic(error)
				fmt.Print(error)
			}
			return tokenString
		} else {
			err = nil
			fmt.Print(err)
			log.Println(err)
			return vo.Mismatch
		}
	}
}

//Signuphandler passes control
func Signupservice(dbuser *vo.DbUser) string {
	var userList vo.Usersobject
	userList.SetUsers(repositories.FindUserByName(dbuser.Username))
	if len(userList.GetUsers()) > 0 {
		log.Println(vo.Exist)
		return vo.Exist
	}
	b := repositories.Insertdb(dbuser)
	if b == true {
		log.Println("Object Inserted")
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": dbuser.Username,
			"password": dbuser.Password,
		})
		tokenString, err := token.SignedString([]byte("secret"))
		if err != nil {
			log.Panic(err)
		}
		log.Println(tokenString)
		dbuser.Token = tokenString
		return tokenString
	} else {
		log.Println(vo.Null)
		return vo.Null
	}
}

//Logouthandler passes control
func Logoutservice(auth string) *jwt.Token {
	if auth != "" {
		bearerauth := strings.Split(auth, " ")
		log.Println("Bearer Authorization: ")
		log.Println(bearerauth)
		if len(bearerauth) == 2 {
			token, error := jwt.Parse(bearerauth[1], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					log.Println("There was an error")
					return nil, fmt.Errorf("There was an error")
				}
				return []byte("secret"), nil
			})
			if error != nil {
				log.Panic(error)
				return nil
			}
			if token.Valid {
				log.Println(token)
				return token
			}
		} else {
			log.Println("null")
			return nil
		}
	}
	return nil
}
