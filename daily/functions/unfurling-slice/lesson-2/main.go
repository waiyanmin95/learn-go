package main

import "fmt"

func main() {
	x := []string{"GG", "EZ", "WP"}
	fmt.Println(x)
}

func foo(s ...string) {
	fmt.Println(s)
}
