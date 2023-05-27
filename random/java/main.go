package main

import "fmt"

func main() {
	var a byte = 255
	a++
	a++

	fmt.Printf("%T\n", a)
	fmt.Println(a)
}
