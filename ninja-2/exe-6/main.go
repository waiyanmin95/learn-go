package main

import "fmt"

const (
	a = 2023 + iota
	b
	c
	d
)

func main() {
	fmt.Println(a, b, c, d)
}
