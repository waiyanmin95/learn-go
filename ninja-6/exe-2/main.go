package main

import "fmt"

func main() {
	aa := []int{2, 4, 6, 8, 10}
	bb := foo(aa...)
	fmt.Println(bb)

	cc := []int{1, 3, 5, 7, 9}
	dd := bar(cc)
	fmt.Println(dd)
}

func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func bar(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
