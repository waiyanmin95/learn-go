package main

import (
	"fmt"
	"log"

	"github.com/ivan-amity/learn-go/exercises/chapter-10-data-encapsulation/setter-method-validation/calendar"
)

const (
	yy = 2022
	mm = 12
	dd = 30
)

func main() {

	// combine
	date := calendar.Date{}

	err := date.SetYear(yy)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetMonth(mm)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetDay(dd)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("SETTER METHOD > YY MM DD:", date.Year, date.Month, date.Day)
}
