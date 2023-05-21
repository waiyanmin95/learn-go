package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("Number of CPUs", runtime.NumCPU())
	fmt.Println("GO Routine", runtime.NumGoroutine())

	var counter int64

	const (
		gs = 100
	)
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			//runtime.Gosched()
			wg.Done()
		}()
		fmt.Println("GO Routine IN LOOP", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("GO Routine", runtime.NumGoroutine())
	fmt.Println("Count", counter)

}
