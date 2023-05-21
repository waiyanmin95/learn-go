package main

import "fmt"

//goland:noinspection ALL
func main() {
	a := 99
	fmt.Println(a)
	fmt.Printf("The type of address %T\n", &a)
	fmt.Println("The address", &a)
	fmt.Println("The value", *&a)
}
