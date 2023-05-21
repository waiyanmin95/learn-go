package main

import "fmt"

func main() {
	aa := make(chan int)

	go func() {
		aa <- 99
	}()

	fmt.Println(<-aa)
}
