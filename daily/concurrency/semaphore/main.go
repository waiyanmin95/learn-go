package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/semaphore"
)

func main() {
	sem := semaphore.NewWeighted(int64(4))
	doneC := make(chan bool, 1)
	totProcess := 10
	ctx := context.Background()
	for i := 1; i <= totProcess; i++ {
		err := sem.Acquire(ctx, 1)
		if err != nil {
			return
		}
		go func(v int) {
			defer sem.Release(int64(1))
			longRunningProcess(v)
			if v == totProcess {
				doneC <- true
			}
		}(i)
	}
	<-doneC
}
func longRunningProcess(taskID int) {
	fmt.Println(
		time.Now().Format("15:04:05"),
		"Running task with ID",
		taskID)
	time.Sleep(2 * time.Second)
}
