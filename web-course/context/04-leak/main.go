package main

import (
	"fmt"
	"time"
)

func main() {
	gen()
	for n := range gen() {
		fmt.Println(n)
		if n == 10 {
			break
		}
	}
	time.Sleep(10 * time.Second)
	fmt.Printf("%T", gen())
}

// gen is a broken generator that will leak a goroutine.
func gen() <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			ch <- n
			n++
		}
	}()
	return ch
}
