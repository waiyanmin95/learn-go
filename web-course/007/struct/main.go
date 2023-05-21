package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

type person struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	buddha := person{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	jesus := person{
		Name:  "Jesus",
		Motto: "Love all!",
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer nf.Close()

	data := []person{buddha, jesus}

	err = tpl.Execute(nf, data)
	if err != nil {
		log.Fatalln(err)
	}
}
