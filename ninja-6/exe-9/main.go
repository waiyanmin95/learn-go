package main

import "fmt"

func main() {
	xxi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	gg := sum(xxi...)
	fmt.Println(gg)

	hh := even(sum, xxi...)
	fmt.Println(hh)

	ii := odd(sum, xxi...)
	fmt.Println(ii)
}

func sum(xi ...int) int {
	aa := 0
	for _, v := range xi {
		aa += v
	}
	return aa
}

func even(f func(xi ...int) int, ev ...int) int {
	var bb []int
	for _, v := range ev {
		if v%2 == 0 {
			bb = append(bb, v)
		}
	}
	return f(bb...)
}

func odd(g func(xi ...int) int, od ...int) int {
	var cc []int
	for _, v := range od {
		if v%2 != 0 {
			cc = append(cc, v)
		}
	}
	return g(cc...)
}
