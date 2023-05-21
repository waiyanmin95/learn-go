package main

import "fmt"

func main() {
	var xi interface{}
	describe(xi)

	xi = "hello"
	describe(xi)

	xi = 44
	describe(xi)
}

func describe(i interface{}) {
	fmt.Printf("(Value: %v, Type: %T)\n", i, i)
}
