//package main
//
//import "fmt"
//
//func main() {
//	fmt.Println(isPalindrome("ARNNRA"))
//}
//
//func isPalindrome(s string) bool {
//	rs := []rune(s)
//	for i := 0; i < len(rs)/2; i++ {
//		if rs[i] != rs[len(rs)-i-1] {
//			return false
//		}
//	}
//	return true
//}

package main

import (
	"fmt"
	"strings"
)

func main() {
	// Get input string
	var input string
	fmt.Print("Enter a string: ")
	fmt.Scan(&input)

	// Remove spaces and make the string lowercase
	input = strings.ToLower(strings.Join(strings.Fields(input), ""))

	// Check if the string is a palindrome
	if input == reverse(input) {
		fmt.Println("The string is a palindrome.")
	} else {
		fmt.Println("The string is not a palindrome.")
	}
}

// Helper function to reverse a string
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
