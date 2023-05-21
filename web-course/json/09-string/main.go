package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	var data string
	rcvd := `"waiyan"`
	fmt.Printf("%T\n", rcvd)
	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%T %v\n", &data, &data)
	fmt.Printf("%T %v\n", data, data)
	fmt.Println(data)
}
