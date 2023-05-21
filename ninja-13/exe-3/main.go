package main

import (
	"fmt"
	"github.com/ivan-amity/learn-go/ninja-13/exe-3/mymath"
)

func main() {
	xxi := gen()
	fmt.Println(xxi)
	for gg, v := range xxi {
		fmt.Println(gg, v)
		fmt.Println(mymath.CenteredAvg(v))
	}
}

func gen() [][]int {
	a := []int{1, 4, 6, 8, 100}
	b := []int{0, 8, 10, 1000}
	c := []int{9000, 4, 10, 8, 6, 12}
	d := []int{123, 744, 140, 200}
	e := [][]int{a, b, c, d}
	return e
}
