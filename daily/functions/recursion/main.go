package main

import "fmt"

func main() {

	fmt.Println(4 * 3 * 2 * 1)
	xx := f(4)
	fmt.Println(xx)

	l := loop(4)
	fmt.Println(l)
}

func f(x int) int {
	if x == 0 {
		return 1
	}
	return x * f(x-1)
}

func loop(f int) int {
	total := 1
	for ; f > 0; f-- {
		total *= f
	}
	return total
}
