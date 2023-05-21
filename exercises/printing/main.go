package main

import "fmt"

func main() {
	fmt.Printf("%%7.3f: %7.3f\n", 12444444.34969)
	fmt.Printf("%%7.3f: %7.3f\n", 124444499.35969)

	fmt.Printf("%12s | %2d\n", "Paper Clips", 4)
	fmt.Printf("%12s | %2d\n", "G", 4)
	fmt.Printf("%12s | %2d\n", "Black", 22)
}
