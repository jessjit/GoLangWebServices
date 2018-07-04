package repositories

import (
	"Loginmodule/data"
	"log"

	"gopkg.in/mgo.v2/bson"
)

func Removedb(id string) {
	var err error
	err = data.Collection.Remove(bson.M{"_id": id})
	if err != nil {
		log.Printf("Could not find Note %s to delete")
	} else {
		log.Println("Object deleted")
	}
}
