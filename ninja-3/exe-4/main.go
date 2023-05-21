package main

import "fmt"

func main() {
	a := 1995
	for {
		if a > 2022 {
			break
		}
		fmt.Println(a)
		a++
	}
}
