package main

import "fmt"

type Wee string

func (m Wee) SayHi() {
	fmt.Println("Hello from:", m)
}

func (l Wee) WithReturn() int {
	return len(l)
}

func main() {
	value := Wee("Merry X'Mas")
	value.SayHi()

	fmt.Println(value.WithReturn())
}
