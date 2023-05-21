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

	fmt.Println(p1)

	m := map[string]person{
		p1.last: p1,
	}
	fmt.Println(m)

	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		fmt.Println(v.fav)
		for i, val := range v.fav {
			fmt.Println(i, val)
		}
	}
}
