package main

import "fmt"

func main() {
	x := make([]int, 10, 10)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 199)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 777)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 8834, 3385, 3333, 4444, 5555, 55553, 12121, 41414, 5551, 56789, 678)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
