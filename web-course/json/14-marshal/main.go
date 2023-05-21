package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type model struct {
	state    bool
	pictures []string
}

// Struct values encode as JSON objects.
// Each exported struct field becomes a member of the object,
// using the field name as the object key
// unless the field is omitted for one of the reasons given below.

func main() {
	m := model{
		state: true,
		pictures: []string{
			"a.jpg",
			"b.jpg",
			"c.jpg",
		},
	}

	fmt.Println(m)

	bs, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	os.Stdout.Write(bs)
}
