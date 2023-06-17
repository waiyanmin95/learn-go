package main

import "fmt"

func main() {
	var number int
	fmt.Printf("Enter element to find: ")
	fmt.Scanln(&number)
	xi := []int{10, 20, 30, 40, 50}

	index := 0
	found := false

	for i := 0; i < len(xi); i++ {
		if xi[i] == number {
			found = true
			index++
			break
		}
	}
	if found == true {
		fmt.Println("Element found at index", index)
	}
}
