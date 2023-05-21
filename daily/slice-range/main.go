package main

import "fmt"

func main() {
	x := []int{3, 5, 7, 9, 11}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])

	//for key, val = range m {
	//	h(key, val)
	//}

	for i, gg := range x {
		fmt.Println(i, gg)
	}
}
