package main

import (
	"fmt"
	"strings"
)

func JoinWithComma(phrases []string) string {
	if len(phrases) == 0 {
		return "Please make an input"
	} else if len(phrases) == 1 {
		return phrases[0]
	} else if len(phrases) == 2 {
		return phrases[0] + " and " + phrases[1]
	} else {
		result := strings.Join(phrases[:len(phrases)-1], ", ")
		result += ", and "
		result += phrases[len(phrases)-1]
		return result
	}
}

func main() {
	//fmt.Println("Happy B-day, GoLand")
	//fmt.Println(JoinWithComma([]string{"apple", "orange", "banana", "grape"}))
	//fmt.Println(JoinWithComma([]string{"apple", "orange"}))
	//fmt.Println(JoinWithComma([]string{"apple"}))
	fmt.Println(JoinWithComma([]string{}))
}
