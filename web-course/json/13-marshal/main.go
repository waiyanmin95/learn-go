package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type model struct {
	State    bool
	Pictures []string
}

func main() {
	m := model{
		State:    true,
		Pictures: nil,
	}

	fmt.Println(m)

	bs, err := json.Marshal(m)
	if err != nil {
		log.Println("Error: ", err)
	}
	fmt.Println(string(bs))
}
