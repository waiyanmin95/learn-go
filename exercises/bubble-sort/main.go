package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println(BubbleSort([]int{22, 10, 6, 2, 1, 5, 8, 3, 4, 7, 9, 100, 30}))
	fmt.Println(SelectionSort([]int{277, 10, 6, 9, 2, 5, 8, 3, 4, 7, 88, 388, 30}))
}

func BubbleSort(xi []int) []int {

	if len(xi) == 1 {
		log.Panicln("Slice should be more than 1 number")
	}

	for j := 0; j < len(xi); j++ {
		for i := 0; i < len(xi)-1; i++ {
			if xi[i] > xi[i+1] {
				xi[i], xi[i+1] = xi[i+1], xi[i]
			}
		}
	}
	return xi
}

func SelectionSort(xi []int) []int {
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
	return xi
}
