package main

import "fmt"

func main() {
	a := foo()
	fmt.Println(a)

	b, bb := bar()
	fmt.Println(b, bb)
}

func foo() int {
	return 5
}

func bar() (int, string) {
	return 88, "HELLO"
}
