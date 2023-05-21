package main

import "fmt"

func main() {
	favSport := "Gaming"
	switch favSport {
	case "Gaming":
		fmt.Println("This is my favourite")
	case "Football":
		fmt.Println("This is my second favourite")
	case "Swimming":
		fmt.Println("This is good for health")
	default:
		fmt.Println("THIS IS DEFAULT")
	}
}
