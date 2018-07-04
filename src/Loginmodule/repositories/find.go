package repositories

import (
	"Loginmodule/RowMappers"
	"Loginmodule/data"
	"Loginmodule/vo"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

//FindUserByName will find user list by username
func FindUserByName(name string) []vo.User {
	fmt.Print("DB running")
	fmt.Println(name)
	dbUsers := []vo.DbUser{}
	//users := []vo.User{}
	conditions := bson.M{"username": name}
	data.Collection.Find(conditions).All(&dbUsers)
	fmt.Print("dbUserList")
	fmt.Println(dbUsers)

	return RowMappers.MapDbUserToUser(dbUsers)
}
