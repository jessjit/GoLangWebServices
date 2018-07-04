package database

import (
	"gopkg.in/mgo.v2"
)

var (
	Session    *mgo.Session
	Collection *mgo.Collection
)

func initializers() {
	Session, err = mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	Session.SetMode(mgo.Monotonic, true)
	Collection = session.DB("userdb").C("usercollection")
}
