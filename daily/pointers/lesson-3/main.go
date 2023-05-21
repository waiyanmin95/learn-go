package main

import "fmt"

func main() {
	x := 0
	fmt.Println("X before", x)
	fmt.Println("X pointer address before", &x)
	foo(&x)
	fmt.Println("X after", x)
	fmt.Println("X pointer address after", &x)
}

func foo(y *int) {
	fmt.Println("Y before", y)
	fmt.Println("Y pointer value before", *y)
	*y = 44
	fmt.Println("Y after", y)
	fmt.Println("Y pointer value after", *y)
}
