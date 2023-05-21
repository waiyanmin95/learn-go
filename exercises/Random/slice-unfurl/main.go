package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5}
	severalInts(xi...)
}

func severalInts(xi ...int) {
	fmt.Println(xi)
}
