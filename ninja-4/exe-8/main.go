package main

import "fmt"

func main() {

	a := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}

	fmt.Println(a)

	for i, v := range a {
		fmt.Printf("This the records for: %v\n", i)
		for j, k := range v {
			fmt.Printf("The Indexex: %v\t %v\t\n", j, k)
		}
	}
}
