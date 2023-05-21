package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://api.github.com/users/waiyanmin95")
	if err != nil {
		log.Fatalf("error :%s", err)
		/*
			log.Printf("error :%s", err)
			os.Exit(1)
		*/
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("error :%s", resp.Status)
	}
	fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))
	fmt.Printf("Server: %s\n", resp.Header.Get("Server"))
	fmt.Printf("Content-Length: %s\n", resp.Header.Get("Content-Length"))

	var r Reply

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		log.Fatalf("Error: Can't Decode %s", err)
	}
	fmt.Println(r)
}

type Reply struct {
	Login           string `json:"login"`
	Name            string `json:"name"`
	Location        string `json:"location"`
	TwitterUsername string `json:"twitter_username"`
}
