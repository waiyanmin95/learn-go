package main

import "fmt"

func main() {
	m := map[string]int{
		"WYM": 27,
		"GUU": 28,
	}

	fmt.Println(m["WYM"])
	fmt.Println(m["GUU"])

	fmt.Println(m["BUU"])

	v, ok := m["SHOO"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["GII"]; ok {
		fmt.Println("THIS IS IF OUTPUT", v)
	}

	m["BLUE"], m["BLACK"] = 34, 44

	for i, v := range m {
		fmt.Println(i, v)
	}

	fmt.Println(m["BLUE"], m["BLACK"])

	delete(m, "GUU")

	fmt.Println(m)
}
