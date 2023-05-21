package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ivan-amity/learn-go/web-course/mongodb/03-post-delete/models"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	// added route plus parameter
	r.HandleFunc("/user/{id}", getUsers).Methods(http.MethodGet)
	r.HandleFunc("/user", createUser).Methods(http.MethodPost)
	r.HandleFunc("/user/{id}", deleteUser).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

func index(w http.ResponseWriter, r *http.Request) {
	s := `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Index</title>
</head>
<body>
<a href="/user/9872309847">GO TO: http://localhost:8080/user/777</a>
</body>
</html>
	`
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pp := params["id"]
	u := models.User{
		Name:   "ABCD",
		Gender: "null",
		Age:    99,
		Id:     pp,
	}

	uj, err := json.Marshal(u)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	// composite literal - type and curly braces
	u := models.User{}

	// encode/decode for sending/receiving JSON to/from a stream
	json.NewDecoder(r.Body).Decode(&u)

	// Change Id
	u.Id = "999"

	// marshal/unmarshal for having JSON assigned to a variable
	uj, err := json.Marshal(u)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	// TODO: write code to delete user
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Write code to delete user\n")
}
