package main

import "fmt"

func main() {
	fmt.Printf("false && false = %t\n", false && false)
	fmt.Printf("true && false = %t\n", true && false)
	fmt.Printf("true && true = %t\n", true && true)

	fmt.Printf("%b & %b == %b\n", 0, 0, 0&0)
	fmt.Printf("%b & %b == %b\n", 0, 1, 0&1)
	fmt.Printf("%b & %b == %b\n", 1, 1, 1&1)

	fmt.Println(170 & 15)
	fmt.Println(10 & 7)
	fmt.Println(100 & 45)

	fmt.Printf("%02b\n", 100)
	fmt.Printf("%02b\n", 45)
}
