package main

import "fmt"

func main() {
	a, err := fmt.Println("HELLO")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
}
