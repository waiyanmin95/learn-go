package main

import (
	"fmt"
)

func main() {
	c1 := struct {
		color   string
		doors   int
		wheels  int
		friends map[string]string
		drinks  []string
	}{
		color:  "White",
		doors:  4,
		wheels: 4,
		friends: map[string]string{
			"F1": "KZY",
			"F2": "JOJO",
			"F3": "TAYOTE",
		},
		drinks: []string{
			"RUM",
			"VODKA",
			"JIN",
		},
	}
	fmt.Println(c1)

	fmt.Println(c1.color, c1.wheels, c1.doors)
	fmt.Println(c1.friends)
	fmt.Println(c1.drinks)

	for i, v := range c1.friends {
		fmt.Println(i, v)
	}

	for ii, vv := range c1.drinks {
		fmt.Println(ii, vv)
	}
}
