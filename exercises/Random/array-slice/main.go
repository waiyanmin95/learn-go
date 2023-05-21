package main

import "fmt"

func main() {
	arr := [5]string{"a", "b", "c", "d", "e"}
	xi := arr[1:3]

	xi = append(xi, "x")
	xi = append(xi, "y", "z")

	for _, v := range xi {
		fmt.Println(v)
	}
}
