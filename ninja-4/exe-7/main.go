package main

import "fmt"

func main() {
	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	xy := [][]string{x, y}
	fmt.Println(xy)

	for i, ok := range xy {
		fmt.Printf("The record is:%v\n", i)
		for v, gg := range ok {
			fmt.Printf("\t Index: %v \tThe value: %v\n", v, gg)
		}
	}
}
