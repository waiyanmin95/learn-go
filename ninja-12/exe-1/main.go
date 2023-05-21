package main

import (
	"fmt"
	"github.com/ivan-amity/learn-go/ninja-12/exe-1/dog"
)

type human struct {
	name string
	age  int
}

func main() {
	nm := human{
		name: "NyoMa",
		age:  dog.Years(10),
	}
	ml := human{
		name: "MelLone",
		age:  dog.Years(7),
	}
	fmt.Println(nm)
	fmt.Println(ml)
}
