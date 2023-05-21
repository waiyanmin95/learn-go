package main

import "fmt"

func main() {
	xi := []int{22, 10, 6, 2, 1, 5, 8, 3, 4, 7, 9, 100, 30}
	xxi := []int{22, 10, 6, 2, 1, 5, 8, 3, 4, 7, 9, 100, 30, 70}

	fmt.Println(xi[0 : len(xi)/2]) // [22 10 6 2 1 5]
	fmt.Println(xi[len(xi)/2:])    // [8 3 4 7 9 100 30]

	//fmt.Println(len(xi))
	fmt.Println(mergeSort(xi))
	fmt.Println(mergeSort(xxi))

}

func mergeSort(s []int) []int {
	if len(s) < 2 {
		return s
	}

	left := s[0 : len(s)/2]
	right := s[len(s)/2:]

	return merge(mergeSort(left), mergeSort(right))
}

func merge(l, r []int) []int {
	results := []int{}

	lcount := 0
	rcount := 0
	a := 0
	b := 0

	for a < len(l) && b < len(r) {
		if l[a] < r[b] {
			results = append(results, l[a])
			a++
		} else {
			results = append(results, r[b])
			b++
		}
	}

	for ; a < len(l); a++ {
		lcount = lcount + 1
		results = append(results, l[a])
	}

	for ; b < len(r); b++ {
		rcount = rcount + 1
		results = append(results, r[b])
	}

	fmt.Println("Right Count", rcount)
	fmt.Println("Left Count", lcount)
	return results
}
