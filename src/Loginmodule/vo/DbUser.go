package vo

import (
	"gopkg.in/mgo.v2/bson"
)

type DbUser struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`
	Username  string        `json:"username" bson:"username"`
	Password  string        `json:"password" bson:"password"`
	Emailid   string        `json:"emailid" bson:"emailid"`
	Firstname string        `json:"firstname" bson:"firstname"`
	Lastname  string        `json:"lastname" bson:"lastname"`
	Age       int           `json:"age" bson:"age"`
	Token     string        `json:"token" bson:"token"`
}
