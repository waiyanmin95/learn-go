package main

import "fmt"

func main() {
	x := 10
	for {
		if x > 10 {
			break
		}
		if x%2 != 0 {
			continue
		}
		fmt.Println(x)
	}
}
