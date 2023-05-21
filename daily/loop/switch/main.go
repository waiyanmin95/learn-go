package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("NOT PRINT")

	case (0 != 0):
		fmt.Println("SHOULD NOT PRINT")

	case (true):
		fmt.Println("SHOULD PRINT")
		fallthrough
	case (100 != 200):
		fmt.Println("SHOULD PRINT, BUT DOES IT PRINT")
		fallthrough
	case (400 != 400):
		fmt.Println("FALLTHROUGH LET's PRINT IT")
		fallthrough
	default:
		fmt.Println("THIS IS A DEFAULT")

	}
}
