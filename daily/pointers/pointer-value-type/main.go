package main

import (
	"fmt"
)

func sliceDelth(ii []int) {
	ii[0] = 99
}

func main() {
	z := 100
	fmt.Println(z)
	x := 42
	fmt.Printf("Value: %v\t\t\t Type: %T\n", &x, &x)

	y := &x
	fmt.Println("Memory Addresses: ", y, "\t\t", &x)
	fmt.Println("Deference address of *y = ", *y, "\t\t\tValue of x :", x)

	*y = 22
	fmt.Println(*y, x)

	xi := []int{1, 2, 3, 4}
	fmt.Printf("%p\n", xi)

	sliceDelth(xi)
	fmt.Printf("%p\n", xi)
	fmt.Println(xi)

	xxi := xi
	fmt.Printf("%p\n", xxi)
}
