package main

import (
	"html/template"
	"log"
	"os"
)

type Invoice struct {
	Name    string
	Paid    bool
	Charges []float64
	Total   float64
}

// check calls log.Fatal on any non-nill error.
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	html, err := template.ParseFiles("bill.html")
	check(err)
	bill := Invoice{
		Name:    "Mary",
		Paid:    true,
		Charges: []float64{23.19, 1.13},
		Total:   44.22,
	}
	err = html.Execute(os.Stdout, bill)
	check(err)
}
