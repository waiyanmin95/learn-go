package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"Guu","Last":"Nandar","Age":28},{"First":"Wai","Last":"Yan","Age":27}]
`
	sb := []byte(s)

	fmt.Println(sb)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", sb)

	var people []person

	err := json.Unmarshal(sb, &people)
	if err != nil {
		fmt.Println("ERROR :", err)
	}
	fmt.Println(people)

	for i, v := range people {
		fmt.Println("No. of people", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}
