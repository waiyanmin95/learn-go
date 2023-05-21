package main

import (
	"fmt"
	"strings"
)

func main() {
	broken := "G# R#CK!"
	replace := strings.NewReplacer("#", "O")
	fix := replace.Replace(broken)

	fmt.Println(fix)
}
