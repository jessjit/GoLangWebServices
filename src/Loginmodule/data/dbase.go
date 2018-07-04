package data

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

var (
	Session    *mgo.Session
	Collection *mgo.Collection
)

func Activatedb(db string, col string) {
	Session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	Collection = Session.DB(db).C(col)
	fmt.Println("DB running")
}
