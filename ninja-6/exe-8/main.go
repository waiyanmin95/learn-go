package main

import "fmt"

func main() {
	ff := foo()
	fmt.Println(foo())
	fmt.Println(ff())
	fmt.Printf("%T\n", ff)
	fmt.Printf("%T\n", foo())
}

func foo() func() int {
	return func() int {
		return 8834
	}
}
