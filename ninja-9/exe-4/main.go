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

	gs := 100
	var mu sync.Mutex
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := incrementer
			v++
			incrementer = v
			fmt.Println(incrementer)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Number of Go Routine", runtime.NumGoroutine())
	fmt.Println("The Final Value", incrementer)
}
