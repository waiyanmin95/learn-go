package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("Number of Go Routine", runtime.NumGoroutine())
	incrementer := 0
	var wg sync.WaitGroup

	const (
		gs = 100
	)
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			fmt.Println(incrementer)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Number of Go Routine", runtime.NumGoroutine())
	fmt.Println("The Final Value", incrementer)
}
