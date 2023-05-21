package main

import "fmt"

func main() {
	//y := foo()

	//fmt.Printf("%T\n", y)
	fmt.Println(foo()())
	fmt.Println(bar())
}

func foo() func() int {
	return func() int {
		return 888
	}
}

func bar() string {
	s := "BLAHBLAHBLUE"
	return s
}
