package main

import (
	"fmt"
	"runtime"
)

func main() {
	s := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			s <- i
		}
		close(s)
	}()

	for v := range s {
		fmt.Println(v)
	}

	fmt.Println("GoRoutines:", runtime.NumGoroutine())
	fmt.Println("EXITING ...")
}
