package main

import (
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"

	"github.com/learn-go/web-course/mongodb/05-mongodb/03-update-user-controller-post/controllers"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/users", uc.GetUser)
	r.POST("/users/:id", uc.CreateUser)
	r.DELETE("/users/:id", uc.DeleteUser)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

func getSession() *mgo.Session {
	s, err := mgo.DialWithTimeout("mongodb://root:secret@localhost:27017", 10*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	return s
}
