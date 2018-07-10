package vo

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	id        bson.ObjectId `json:"id"`
	username  string        `json:"username"`
	password  string        `json:"password"`
	emailid   string        `json:"emailid"`
	firstname string        `json:"firstname"`
	lastname  string        `json:"lastname"`
	age       int           `json:"age"`
	token     string        `json:"token"`
}

func CreateUserObj(name string, word string) (usr User) {
	usr.username = name
	usr.password = word
	return usr
}

func (usr User) Getuserid() bson.ObjectId {
	return usr.id
}

func (usr User) Getusername() string {
	return usr.username
}

func (usr User) Getuserpassword() string {
	return usr.password
}

func (usr User) Getemailid() string {
	return usr.emailid
}

func (usr User) Getfirstname() string {
	return usr.firstname
}

func (usr User) Getlastname() string {
	return usr.lastname
}

func (usr User) Getuserage() int {
	return usr.age
}

func (usr User) Gettoken() string {
	return usr.token
}

func (usr *User) Setid(id bson.ObjectId) {
	usr.id = id
}

func (usr *User) Setusername(username string) {
	usr.username = username
}

func (usr *User) Setpassword(password string) {
	usr.password = password
}

func (usr *User) Setemailid(emailid string) {
	usr.emailid = emailid
}

func (usr *User) Setfirstname(firstname string) {
	usr.firstname = firstname
}

func (usr *User) Setlastname(lastname string) {
	usr.lastname = lastname
}

func (usr *User) Setuserage(age int) {
	usr.age = age
}

func (usr *User) Settoken(token string) {
	usr.token = token
}
