package main

import "fmt"

func main() {
	x := []int{3, 5, 7, 9, 11}
	fmt.Println(x)
	fmt.Println(x[0])

	fmt.Println(x[:])

	fmt.Println(x[1:])

	// 2nd place not including the item
	fmt.Println(x[1:3])

	for i, v := range x {
		fmt.Println(i, v)
	}

	for i := 0; i <= 4; i++ {
		fmt.Println(i, x[i])
	}
}
