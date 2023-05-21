package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{88, 33, 22, 83, 1, 6, 99, 77, 123, 8, 0}
	xs := []string{"Ball", "Gall", "Yee", "Blue", "Black", "White", "Green", "Yell"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}
