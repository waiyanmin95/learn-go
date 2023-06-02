package main

import (
	"fmt"
	"strings"
)

func main() {
	c := 'B'
	a := "GG"
	b := "GG"

	fmt.Println("Check Equal: ", a == b)
	fmt.Printf("Value of C: %d\n", c)
	fmt.Println(strings.Compare(a, b))
}
