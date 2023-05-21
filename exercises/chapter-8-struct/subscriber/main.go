package main

import "fmt"

type Subscriber struct {
	name   string
	rate   float64
	active bool
}

func main() {
	PrintInfo(Default("Wai Yan"))

	fmt.Println("-------")
	sub1 := Default("GUU")
	sub1.rate = 3.99
	PrintInfo(sub1)
}

func PrintInfo(s Subscriber) {
	fmt.Println(s.name)
	fmt.Println(s.rate)
	fmt.Println(s.active)
}

func Default(s string) Subscriber {
	aa := Subscriber{
		name:   s,
		rate:   4.99,
		active: true,
	}
	return aa
}
