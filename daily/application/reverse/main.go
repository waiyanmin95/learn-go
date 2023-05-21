package main

import (
	"fmt"
	"sort"
)

func main() {
	si := []int{5, 7, 4, 99, 22, 100, 101, 33, 22, 13}
	sort.Sort(sort.Reverse(sort.IntSlice(si)))
	fmt.Println(si)
}
