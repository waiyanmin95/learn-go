package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/ivan-amity/learn-go/web-course/mongodb/04-controllers/models"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

// GetUser Methods have to be capitalized to be exported, eg, GetUser and not getUser
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "Ivan",
		Gender: "male",
		Age:    22,
		Id:     p.ByName("id"),
	}
	uj, err := json.Marshal(u)
	if err != nil {
		log.Println("Error", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200 return
	fmt.Fprintf(w, "%s\n", uj)
}

// CreateUser Methods have to be capitalized to be exported, eg, CreateUser and not createUser
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	u.Id = "999"

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201 return
	fmt.Fprintf(w, "%s\n", uj)
}

// DeleteUser Methods have to be capitalized to be exported, eg, DeleteUser and not deleteUser
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// TODO: only write code to delete user
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Write code to delete user\n")
}
