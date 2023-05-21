package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "Guu",
		Last:  "Nandar",
		Age:   28,
	}

	p2 := person{
		First: "Wai",
		Last:  "Yan",
		Age:   27,
	}

	people := []person{
		p1,
		p2,
	}

	fmt.Println(people)

	sb, err := json.Marshal(people)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(sb)
	fmt.Println(string(sb))
	r := string(sb)
	for i, v := range r {
		fmt.Println(i, v)
	}
	fmt.Println(r)
}
