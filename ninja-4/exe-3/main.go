package main

import "fmt"

func main() {
	aa := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(aa)

	fmt.Println(aa[:5])

	fmt.Println(aa[5:])

	fmt.Println(aa[2:7])

	fmt.Println(aa[1:6])
}
