package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

type scores struct {
	S1 int `template:"bold"`
	S2 int `template:"bold"`
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	s := scores{
		S1: 10,
		S2: 8,
	}

	err := tpl.Execute(os.Stdout, s)
	if err != nil {
		log.Fatalln(err)
	}
}
