package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	banner("Go", 6)

	banner("G😀", 6)

	emoji := "G😀"
	fmt.Println(emoji, len(emoji))
	fmt.Printf("%5s %#v\n", emoji, emoji)

	for i, r := range emoji {
		fmt.Println(i, r)
		if i == 0 {
			fmt.Printf("%c of Type %T\n", r, r)
			// rune ( int32 )
		}
	}

	// byte ( unit8 )
	b := emoji[0]
	fmt.Printf("%c of Type %T\n", b, b)
}

func banner(text string, width int) {
	padding := (width - utf8.RuneCountInString(text)) / 2
	//padding := (width - len(text)) / 2 // BUG: len is in the byte

	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}
	fmt.Println(text)
	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}
