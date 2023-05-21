package main

import "fmt"

func main() {
	one()
}

func one() {
	defer fmt.Println("Deferred in one()")
	two()
}

func two() {
	defer fmt.Println("Deferred in two()")
	three()
	//panic("Let's see what's ben deferred!")
}

func three() {
	defer fmt.Println("Deferred in three()")
	panic("Let's see what's ben deferred!")
}
