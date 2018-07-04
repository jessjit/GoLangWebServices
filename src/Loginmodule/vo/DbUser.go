package vo

import (
	jwt "github.com/dgrijalva/jwt-go"
	"gopkg.in/mgo.v2/bson"
)

type DbUser struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`
	Username  string        `json:"username" bson:"username"`
	Password  string        `json:"password"`
	Emailid   string        `json:"emailid"`
	Firstname string        `json:"firstname"`
	Lastname  string        `json:"lastname"`
	Age       int           `json:"age"`
	Token     *jwt.Token    `json:"token"`
}
