package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}

	fmt.Println(m)

	m["WYM"] = []string{"Guu", "Nandar", "Chit"}
	fmt.Println("Before Delete", m)

	delete(m, "bond_james")

	fmt.Println(m)

	for i, e := range m {
		fmt.Println(i)
		for j, v := range e {
			fmt.Println(j, v)
		}
	}
}
