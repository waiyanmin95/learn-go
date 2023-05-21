package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	fmt.Println("CPUS", runtime.NumCPU())
	fmt.Println("Go Routines", runtime.NumGoroutine())

	wg.Add(2)
	go func() {
		fmt.Println("This is ONEEEE")
		wg.Done()
	}()

	go func() {
		fmt.Println("THIS IS TWOOOOO")
		fmt.Println("THIS IS TWOOOOO")
		wg.Done()
	}()

	fmt.Println("Go Routines", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Go Routines", runtime.NumGoroutine())
}
