package services

import (
	"fmt"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

func Logoutservice(auth string) *jwt.Token {
	fmt.Println("auth")
	fmt.Println(auth)
	if auth != "" {
		bearerauth := strings.Split(auth, " ")
		fmt.Println("bearerauth")
		fmt.Println(bearerauth)
		if len(bearerauth) == 2 {
			token, error := jwt.Parse(bearerauth[1], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					fmt.Println("There was an error")
					return nil, fmt.Errorf("There was an error")
				}
				return []byte("secret"), nil
			})
			if error != nil {
				panic(error)
				return nil
			}
			fmt.Println(token)
			//validate the token
			if token.Valid {
				return token
			}
		} else {
			return nil
		}
	}
	return nil
}
