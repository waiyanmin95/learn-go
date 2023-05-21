package main

import "fmt"

func main() {
	x := []int{3, 5, 7, 9, 11}
	fmt.Println(x)

	x = append(x, 2, 4, 8, 10, 12)
	fmt.Println("Append", x)

	y := []int{888, 111, 222, 333, 444}
	fmt.Println(y)
	x = append(x, y...)
	fmt.Println(x)

	// deleting the slice
	x = append(x[:3], x[7:]...)
	fmt.Println(x)
}
