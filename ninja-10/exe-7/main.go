package main

import (
	"fmt"
	"runtime"
)

func main() {
	c := make(chan int)

	goroutines := 10
	for k := 0; k < goroutines; k++ {
		go func() {
			for i := 0; i < goroutines; i++ {
				c <- i
			}
		}()

	}

	for j := 0; j < 100; j++ {
		fmt.Println("GoRoutines:", runtime.NumGoroutine())
		fmt.Println(j, <-c)
	}
	fmt.Println("EXITING ....")
}
