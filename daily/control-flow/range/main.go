package main

import "fmt"

func main() {
	xi := []int{10, 20, 30, 40, 50}

	for i, v := range xi {
		fmt.Println("Index:", i, "Value:", v)
		v *= 2
	}
	fmt.Println("Raw:", xi)
	fmt.Println("Raw:", xi[0])
	fmt.Println("--------------------------")

	for i, v := range xi {
		fmt.Println("Index:", i, "Value:", v)
		xi[i] *= 2
	}
	fmt.Println("Manipulated:", xi)
	fmt.Println("Manipulated:", xi[0])
}
