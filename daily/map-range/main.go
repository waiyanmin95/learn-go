package main

import "fmt"

func main() {
	xi := []int{33, 44, 55, 66, 77, 88}

	a, b := 42, 43

	fmt.Println(a, b)

	for o, v := range xi {
		fmt.Println(o, v)
	}
}
