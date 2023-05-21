package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretagent struct {
	person
	cankill bool
}

func main() {
	sa1 := secretagent{
		person: person{
			first: "WAI YAN",
			last:  "MIN",
			age:   27,
		},
		cankill: true,
	}

	fmt.Println(sa1.cankill, sa1.first, sa1.last, sa1.age, sa1.person)
}
