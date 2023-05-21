package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, b, c, d, e int
	fmt.Scan(&a, &b, &c, &d, &e)

	xi := make([]int, 0)
	xi = append(xi, a, b, c, d, e)

	sort.Ints(xi)
	min := 0
	max := 0

	for i := 0; i < len(xi); i++ {
		if i < len(xi)-1 {
			min += xi[i]
		}
		if i > 0 {
			max += xi[i]
		}
	}
	fmt.Println(min, max)
}
