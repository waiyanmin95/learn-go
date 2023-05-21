package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(&a) // & gives you the address
	fmt.Printf("%T\n", &a)

	b := &a
	fmt.Println(b)
	fmt.Println(*b) // * gives you the value stored at an address when you have the address
	fmt.Printf("%T\n", *b)

	fmt.Println(*&a)

	*b = 77
	fmt.Println(a)
}
