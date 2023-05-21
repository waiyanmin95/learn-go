package main

import "fmt"

func main() {
	// for init; condition; post {}
	for i := 0; i <= 10; i++ {
		fmt.Printf("The Outter Loop: %d\n", i)
		for j := 1; j <= 3; j++ {
			fmt.Printf("\t\t The Inner Loop: %d\n", j)
		}

	}

}
