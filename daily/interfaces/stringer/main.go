package main

import (
	"fmt"
	"log"
	"strconv"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprint("The name of book is ", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("The number is ", strconv.Itoa(int(c)))
}

func logInfo(s fmt.Stringer) {
	log.Println("LOG FROM LESSON", s.String())
}

func main() {
	x := book{title: "Holy Moly"}

	var y count = 42
	fmt.Println(x)
	fmt.Println(y)

	logInfo(x)
	logInfo(y)
}
