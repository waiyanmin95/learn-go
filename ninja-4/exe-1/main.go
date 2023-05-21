package main

import "fmt"

func main() {
	aa := [5]int{100, 200, 300, 400, 500}
	fmt.Println(aa)

	for i, v := range aa {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", aa)
}
