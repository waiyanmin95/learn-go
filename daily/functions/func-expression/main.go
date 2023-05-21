package main

import "fmt"

func main() {
	fmt.Println("HELLO WORLD")

	func() {
		fmt.Println("Anonymous")
	}()

	f := func() {
		fmt.Println("Function Expression")
	}
	f()

	g := func(x string) {
		fmt.Println("Function Expression with Anonymous", x)
	}
	g("EZ")
}
