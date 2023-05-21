package main

import "fmt"

func main() {
	var x int
	x++
	fmt.Println(x)
	i := c()
	fmt.Println(i)
}

func c() (i int) {
	defer func() { i++ }()
	return 1
}
