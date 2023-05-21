package main

import "fmt"

func main() {
	a := []string{"RUM", "VODKA", "RED", "BLACK"}
	fmt.Println(a)

	b := []string{"WHISKY", "JIN", "BLUE", "GREEN"}
	fmt.Println(b)

	m := [][]string{a, b}
	fmt.Println(m)
}
