package main

import "fmt"

func main() {
	gg := foo()
	fmt.Printf("THIS is the Type GG, %T\n", gg)
	fmt.Println(gg())
	fmt.Println(gg())
	fmt.Println(gg())
	fmt.Println(gg())
	fmt.Println(gg())

	hh := foo()
	fmt.Printf("THIS is the Type HH, %T\n", hh)
	fmt.Println(hh())
	fmt.Println(hh())
	fmt.Println(hh())
}

func foo() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
