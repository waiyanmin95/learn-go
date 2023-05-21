package main

import "fmt"

func main() {
	fmt.Println("Two + Three :", mySum(2, 3))
	fmt.Println("Five + Nine :", mySum(5, 9))
	fmt.Println("Ten + One :", mySum(10, 1))
}

func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
