package data

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
)

var (
	session    *mgo.Session
	Collection *mgo.Collection
)

func Activatedb(db string, col string) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	Collection = session.DB(db).C(col)
	log.Println("Database active and running")
	fmt.Println("Database actice and running")
}
