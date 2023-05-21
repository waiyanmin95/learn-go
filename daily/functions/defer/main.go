package main

import "fmt"

func main() {
	foo()
	defer bar()
	woo()
}

func foo() {
	fmt.Println("THIS IS FOO")
}

func bar() {
	fmt.Println("THIS IS BAR")
}

func woo() {
	fmt.Println("THIS IS WOO")
}
