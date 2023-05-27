package main

import "fmt"

func main() {
	n := "EGGE"
	switch n {
	case "WYM":
		fmt.Println("THIS IS WYM and SHOULD NOT PRINT")
	case "GOLAND":
		fmt.Println("THIS SHOULD NOT PRINT")
	case "IDOL", "EGGE":
		fmt.Println("THIS SHOULD PRINT")
		fallthrough
	default:
		fmt.Println("THIS IS DEFAULT")
	}
}
