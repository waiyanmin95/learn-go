package main

import (
	"fmt"
	"github.com/ivan-amity/learn-go/ninja-13/exe-2/quote"
	"github.com/ivan-amity/learn-go/ninja-13/exe-2/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
