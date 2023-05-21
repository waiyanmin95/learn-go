package main

import "fmt"

func main() {
	x := []int{3, 5, 7, 9, 11}
	fmt.Println(x)

	x = append(x, 2, 4, 8, 10, 12)
	fmt.Println(x)

	y := []int{888, 111, 222, 333, 444}
	fmt.Println(y)
	x = append(x, y...)
	fmt.Println(x)
}

//func one() {
//	x := []int{1, 2, 3}
//	y := []int{4, 5, 6}
//	x = append(x, y...)
//	fmt.Println(x)
//
//}
