package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	s := "Biology is the study of living things."

	s64 := base64.StdEncoding.EncodeToString([]byte(s))

	d, err := base64.StdEncoding.DecodeString(s64)
	if err != nil {
		log.Fatalln("This is ecology", err)
		return
	}

	fmt.Println(d)
	fmt.Println("Length", len(d))
	fmt.Println(string(d))
}
