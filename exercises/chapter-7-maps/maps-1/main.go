package main

import "fmt"

func main() {
	ranks := make(map[string]int)

	people := map[string]string{"waiyan": "hello", "guu": "world"}

	fmt.Printf("%#v\n", people["blah"])

	fmt.Println(people)

	counter := make(map[string]int)

	counter["a"]++
	counter["a"]++
	counter["b"]++

	fmt.Println("Counter:", counter["a"], counter["b"], counter["c"])

	ranks["gold"] = 1
	ranks["silver"] = 2
	ranks["bronze"] = 3

	fmt.Println(ranks)
	//fmt.Println(ranks["gold"])
	//fmt.Println(ranks["jade"])

	ranks["jade"] = 5
	ranks["car"] = 2
	ranks["apple"] = 33
	ranks["zoo"] = 1
	fmt.Println(ranks)
}
