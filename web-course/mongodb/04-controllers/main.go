package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/ivan-amity/learn-go/web-course/mongodb/04-controllers/controllers"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController()
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
