package main

import "fmt"

type person struct {
	first string
	last  string
	fav   []string
}

func main() {
	p1 := person{
		first: "Wai Yan",
		last:  "Min",
		fav: []string{
			"Chocolate",
			"Strawbery",
			"Vanila",
		},
	}
	fmt.Println(p1.first, p1.last, p1.fav)

	for i, v := range p1.fav {
		fmt.Println(i, v)
	}
}
