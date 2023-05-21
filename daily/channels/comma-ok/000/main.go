package main

import (
	"fmt"
)

func main() {
	c := make(chan string)
	go func() {
		c <- "HELLO"
	}()

	v, ok := <-c

	fmt.Println(v, ok)
}
