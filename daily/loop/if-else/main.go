package main

import "fmt"

func main() {
	x := 43
	if x == 40 {
		fmt.Println("Value is 40")
	} else if x == 42 {
		fmt.Println("Value is 42")
	} else if x == 43 {
		fmt.Println("Value is 43")
	} else {
		fmt.Println("Value is not 40, 42 or 43")
	}
}
