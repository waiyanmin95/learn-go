package main

import (
	"fmt"
	"log"
)

func main() {
	xi := []int{22, 10, 6, 2, 1, 5, 8, 3, 4, 7, 9, 100, 30}

	if len(xi) == 1 {
		log.Panicln("Slice should be more than 1 number")
	}

	for i := 0; i < len(xi); i++ {
		for v := 0; v < len(xi); v++ {
			if xi[i] < xi[v] {
				xi[i], xi[v] = xi[v], xi[i]
			}
		}
	}
	fmt.Println("Selection Sort", xi)
}
