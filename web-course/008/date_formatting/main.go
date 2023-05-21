package main

import (
	"html/template"
	"log"
	"os"
	"time"
)

var tpl *template.Template

var fm = template.FuncMap{ // Function Mapping
	"formatDMY": dayMonthYear, // Template Function Name: Go Function
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func dayMonthYear(t time.Time) string {
	//return t.Format("02-01-2006 03:04:05PM")
	return t.Format(time.Kitchen)
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
