package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ivan-amity/learn-go/web-course/mongodb/02-json/models"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	// added route plus parameter
	r.HandleFunc("/user/{id}", getUsers).Methods(http.MethodPost, http.MethodGet)
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
		Name:   "Ivan",
		Gender: "male",
		Age:    22,
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
