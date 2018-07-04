package RowMappers

import (
	"Loginmodule/vo"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

func MapDbUserToUser(dbUser []vo.DbUser) []vo.User {

	var userList []vo.User

	if dbUser != nil && len(dbUser) > 0 {
		for i := 0; i < len(dbUser); i++ {
			var userTemp vo.User
			fmt.Print("i")
			fmt.Println(i)
			if dbUser[i].Id != bson.ObjectId(0) {
				userTemp.Setid(dbUser[i].Id)
			}
			if dbUser[i].Username != "" {
				userTemp.Setusername(dbUser[i].Username)
			}
			if dbUser[i].Password != "" {
				userTemp.Setpassword(dbUser[i].Password)
			}
			if dbUser[i].Emailid != "" {
				userTemp.Setemailid(dbUser[i].Emailid)
			}
			if dbUser[i].Firstname != "" {
				userTemp.Setfirstname(dbUser[i].Firstname)
			}
			if dbUser[i].Lastname != "" {
				userTemp.Setlastname(dbUser[i].Firstname)
			}
			if dbUser[i].Age != 0 {
				userTemp.Setuserage(dbUser[i].Age)
			}
			fmt.Print("userTemp")
			fmt.Println(userTemp)
			userList = append(userList, userTemp)
		}
	}

	return userList
}
