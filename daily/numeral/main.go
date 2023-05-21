package main

import "fmt"

const (
	a int     = 32
	b float64 = 55.43
	c string  = "WAI"
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
