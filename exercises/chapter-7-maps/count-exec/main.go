package main

import "fmt"

func main() {
	data := []string{"a", "c", "e", "a", "e"}
	counts := make(map[string]int)

	for _, i := range data {
		counts[i]++
	}

	letters := []string{"a", "b", "c", "d", "e"}
	for _, l := range letters {
		count, ok := counts[l]
		if !ok {
			fmt.Printf("%s: not found\n", l)
		} else {
			fmt.Printf("%s: %d\n", l, count)
		}
	}

}
