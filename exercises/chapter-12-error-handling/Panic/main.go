package main

import "fmt"

func main() {
	notes := []string{"do", "yay", "mi", "fa", "so", "la", "ti", "do"}

	for i := 0; i <= len(notes); i++ {
		fmt.Println(i, notes[i])
	}
}
