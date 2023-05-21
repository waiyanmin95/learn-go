package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"

	"github.com/ivan-amity/learn-go/web-course/mongodb/05-mongodb/01-update-user-controller/controllers"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://admin:secret@localhost:27017")

	// Check if connection error, is mongo running?
	if err != nil {
		log.Fatalln("Can't connect to MongoDB", err)
	}
	return s
}
