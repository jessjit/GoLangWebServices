package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	session    *mgo.Session
	collection *mgo.Collection
)

type Dogprofile struct {
	Id    bson.ObjectId `bson:"_id" json:"id"`
	Name  string        `json:"name"`
	Breed string        `json:"breed"`
}

type Dogobject struct {
	Dog Dogprofile `json:"dog"`
}
type Dogsobject struct {
	Dogs []Dogprofile `json:"dogs"`
}

func Getdogshandler(w http.ResponseWriter, r *http.Request) {
	var dcol []Dogprofile
	iter := collection.Find(nil).Iter()
	result := Dogprofile{}
	for iter.Next(&result) {
		dcol = append(dcol, result)
	}
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(Dogsobject{Dogs: dcol})
	if err != nil {
		panic(err)
	}
	w.Write(j)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/dogs", Getdogshandler).Methods("GET")
	//router.HandleFunc("/api/dogs", Createdoghandler).Methods("POST")
	//router.HandleFunc("/dogs/{id}", Updatedoghandler).Methods("PUT")
	//router.HandleFunc("/dogs/{id}", Deletedoghandler).Methods("DELETE")
	http.Handle("/api/", router)

	log.Println("Starting the mongo DB session:")
	var err error
	session, err = mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	collection = session.DB("dogsdatabase").C("dcollections")

	log.Println("Listening on 8080")
	http.ListenAndServe(":8000", nil)
}
