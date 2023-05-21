package main

import (
	"fmt"
	"log"
)

func Social() error {
	defer fmt.Println("Good Bye!!!")
	fmt.Println("Hello world")
	return fmt.Errorf("I don't want to talk.")
	fmt.Println("First One")
	fmt.Println("Second One")
	return nil
}

func main() {
	fmt.Println("This is from the main")
	err := Social()
	if err != nil {
		log.Fatal(err)
	}
}
