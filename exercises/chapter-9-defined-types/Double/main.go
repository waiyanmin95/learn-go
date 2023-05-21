package main

import "fmt"

type Num int

func (n *Num) Display() {
	fmt.Println(*n)
}

func (n *Num) Double() {
	*n *= 2
}

func main() {
	number := Num(34)
	number.Double()
	number.Display()
}
