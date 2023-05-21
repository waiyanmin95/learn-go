package main

import "fmt"

func main() {
	gg := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sum(gg...)
	fmt.Println(s)

	z := even(sum, gg...)
	fmt.Println(z)

	yy := odd(sum, gg...)
	fmt.Println(yy)
}

func sum(xi ...int) int {
	fmt.Printf("%T\n", xi)
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func even(f func(xi ...int) int, aa ...int) int {
	var bb []int
	for _, v := range aa {
		if v%2 == 0 {
			bb = append(bb, v)
		}
	}
	return f(bb...)
}

func odd(f func(xi ...int) int, cc ...int) int {
	var dd []int
	for _, v := range cc {
		if v%2 != 0 {
			dd = append(dd, v)
		}
	}
	return f(dd...)
}
