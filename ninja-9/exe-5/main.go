package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("Number of Go Routine", runtime.NumGoroutine())
	var wg sync.WaitGroup

	gs := 100
	var atm int32
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt32(&atm, 1)
			r := atomic.LoadInt32(&atm)
			//runtime.Gosched()
			fmt.Println(r)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Number of Go Routine", runtime.NumGoroutine())
	fmt.Println("This is the final value", atm)
}
