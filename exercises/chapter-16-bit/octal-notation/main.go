package main

import "fmt"

func main() {
	for i := 0; i <= 19; i++ {
		fmt.Printf("%3d: %04o\n", i, i) //base10 //base8
	}
}
