package main

import "fmt"

func main() {
	aa := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println(aa)

	for i, v := range aa {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", aa)
}
