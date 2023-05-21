package main

import (
	"fmt"
	"log"

	getcalendar "github.com/ivan-amity/learn-go/exercises/chapter-10-data-encapsulation/getter-method/calendar"
)

const (
	tt = "Takal top di nae ka"
	yy = 2022
	mm = 12
	dd = 30
)

func main() {

	// combine
	date := getcalendar.Date{}

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

	event := getcalendar.Event{}
	err = event.SetTitle(tt)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("GETTER METHOD", event.Title(), date.Year(), date.Month(), date.Day())
}
