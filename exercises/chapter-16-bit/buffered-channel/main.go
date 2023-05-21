package main

import (
	"fmt"
	"time"
)

func sendLetters(c chan string) {
	time.Sleep(time.Second * 1)
	c <- "a"
	time.Sleep(time.Second * 1)
	c <- "b"
	time.Sleep(time.Second * 1)
	c <- "c"
	time.Sleep(time.Second * 1)
	c <- "d"
}

func main() {
	fmt.Println(time.Now())
	channel := make(chan string, 3e)
	go sendLetters(channel)
	time.Sleep(time.Second * 5)
	fmt.Println(<-channel, time.Now())
	fmt.Println(<-channel, time.Now())
	fmt.Println(<-channel, time.Now())
	fmt.Println(<-channel, time.Now())
	fmt.Println(time.Now())
}
