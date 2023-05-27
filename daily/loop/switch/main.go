package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("NOT PRINT")
	case true:
		fmt.Println("SHOULD PRINT")
		fallthrough
	default:
		fmt.Println("THIS IS A DEFAULT")
	}
}
