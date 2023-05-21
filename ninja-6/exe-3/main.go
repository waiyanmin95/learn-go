package main

import "fmt"

func main() {
	foo()
}

func foo() {
	fmt.Println("THIS ONE GO FIRST")
	defer fmt.Println("ThIS ONE GO SECOND, but SKIPPED with defer")
	fmt.Println("THIS ONE GO THIRD")
}
