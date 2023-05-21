package main

import "fmt"

func main() {
	aa := map[string]int{
		"WYM": 44,
		"GUU": 33,
		"FOO": 55,
		"BAR": 77,
	}

	fmt.Println(aa["WYM"])

	delete(aa, "WYM")
	fmt.Println(aa)

	if v, gg := aa["FOO"]; gg {
		fmt.Println("The VALUE:", v)
		delete(aa, "BAR")
		fmt.Println(aa)
		delete(aa, "GUU")
		fmt.Println(aa)
		delete(aa, "FOO")
		fmt.Println(aa)
	}

}
