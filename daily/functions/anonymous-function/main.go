package main

import "fmt"

func main() {
	foo()

	func() {
		fmt.Println("THIS IS ANONYMOUS FUNCTION")
	}()

	func(x int) {
		fmt.Println("THIS ANONYMOUS FUNCTION", x)
	}(25)

	func(y string) {
		fmt.Println("THIS ANONYMOUS FUNCTION", y)
	}("LUGYAN GYI")

	func(z bool) {
		fmt.Println("THIS ANONYMOUS FUNCTION", z)
	}(true)
}

func foo() {
	fmt.Println("HELLO, THIS IS FOO")
}
