package main

import "fmt"

func main() {
	a := increase()
	b := increase()
	
	fmt.Println("THE VALUEs (A) are: ")
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	fmt.Println("THE VALUEs (B) are: ")
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}

func increase() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
