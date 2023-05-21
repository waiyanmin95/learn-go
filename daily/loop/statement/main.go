package main

import "fmt"

func main() {
	i := 0
	for {
		if i > 9 {
			fmt.Println("Bye")
			break
		}
		fmt.Println(i)
		i++
	}
}
