package main

import "fmt"

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Printf("%d\n", i)

		for y := 0; y < 3; y++ {
			fmt.Printf("\t%#U\n", i)
		}

	}
}
