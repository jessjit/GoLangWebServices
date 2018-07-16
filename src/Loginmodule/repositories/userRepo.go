package repositories

import (
	"Loginmodule/RowMappers"
	"Loginmodule/data"
	"Loginmodule/vo"
	"log"

	"github.com/dgrijalva/jwt-go"

	"gopkg.in/mgo.v2/bson"
)

//FindUserByName will find user list by username
func FindUserByName(name string) []vo.User {
	var dbUsers = []vo.DbUser{}
	conditions := bson.M{"username": name}
	err := data.Collection.Find(conditions).All(&dbUsers)
	if err != nil {
		panic(err)
	}
	return RowMappers.MapDbUserToUser(dbUsers)
}

//Find user by the token
func FindUserByToken(token *jwt.Token) []vo.User {
	tokendbUsers := []vo.DbUser{}
	tokenconditions := bson.M{"token": *token}
	data.Collection.Find(tokenconditions).All(&tokendbUsers)
	return RowMappers.MapDbUserToUser(tokendbUsers)
}

//Insert user to DB
func Insertdb(dbuser *vo.DbUser) (val bool) {
	err := data.Collection.Insert(dbuser)
	if err != nil {
		panic(err)
		return false
	} else {
		return true
	}
}

//Removes Users token from DB
func Removedb(id bson.ObjectId) {
	var err error
	err = data.Collection.Update(bson.M{"_id": id}, bson.M{"$set": bson.M{"token": nil}})
	if err != nil {
		log.Printf("Could not find Note %s to delete")
	} else {
		log.Println("Token deleted")
	}
}

//Set
func Settokendb(name string, token string) bool {
	var err error
	err = data.Collection.Update(bson.M{"username": name}, bson.M{"$set": bson.M{"token": token}})
	if err != nil {
		return false
	} else {
		return true
	}
}
