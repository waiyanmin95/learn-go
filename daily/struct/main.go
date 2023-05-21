package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "Wai Yan",
		last:  "Min",
		age:   27,
	}

	p2 := person{
		first: "Guu",
		last:  "Nandar Chit",
		age:   29,
	}

	fmt.Println(p1.age, p1.last, p1.first)
	fmt.Println(p2.first, p2.age, p2.last)

}
