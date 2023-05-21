package main

import "fmt"

type People struct {
	age int
}

func main() {
	x := People{
		age: 27,
	}
	x.age = 33

	var pointer *People = &x

	fmt.Println((*pointer).age)
}
