package main

import "fmt"

func main() {
	xi := []int{2, 4, 6, 8, 10, 12, 14}
	x := foo(xi...)
	fmt.Println("Total values is =", x)

}

func foo(s ...int) int {
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	sum := 0
	for i, v := range s {
		sum += v
		fmt.Println(i, v, sum)
	}
	return sum
}
