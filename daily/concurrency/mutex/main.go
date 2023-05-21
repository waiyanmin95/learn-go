package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Number of CPUs", runtime.NumCPU())
	fmt.Println("GO Routine", runtime.NumGoroutine())
	counter := 0

	const (
		gs = 100
	)
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			//time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("GO Routine IN LOOP", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("GO Routine", runtime.NumGoroutine())
	fmt.Println("Count", counter)

}
