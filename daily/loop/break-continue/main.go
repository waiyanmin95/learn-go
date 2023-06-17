package main

import "fmt"

func main() {
	x := 10
	for {
		if x%2 == 0 {
			continue
		}
		if x > 10 {
			break
		}
		fmt.Println(x)
	}
}
