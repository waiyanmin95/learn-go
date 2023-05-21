package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("view.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

func executeTemplate(text string, data interface{}) {
	tmpl, err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, data)
}
func main() {
	//text := "Here's my template!\n"
	//tmpl, err := template.New("test").Parse(text)
	//check(err)
	//err = tmpl.Execute(os.Stdout, nil)
	//check(err)

	//http.HandleFunc("/guestbook", viewHandler)
	//err := http.ListenAndServe("localhost:8080", nil)
	//check(err)

	//executeTemplate("Dot is: {{.}}!\n", "ABC")
	//executeTemplate("Dot is: {{.}}!\n", 123.5)
	//
	//executeTemplate("Start {{if .}} Dot is true!{{end}} finish\n", true)
	//executeTemplate("Start {{if .}} Dot is true!{{end}} finish\n", false)

	//templateText := "Before Loop: {{.}}\n{{range .}}In Loop: {{.}}\n{{end}}After Loop: {{.}}"
	//executeTemplate(templateText, []string{"Do", "Yay", "Mi", "Fa"})

	//templateText := "Prices:\n{{range .}}Dollar > ${{.}}\n{{end}}"
	//executeTemplate(templateText, []float64{1.25, 3.2, 4.8})
	//executeTemplate(templateText, []float64{})
	//executeTemplate(templateText, nil)

	//type Part struct {
	//	Name  string
	//	Count int
	//}
	//templateText := "Name: {{.Name}}\nCount: {{.Count}}\n"
	//executeTemplate(templateText, Part{
	//	Name:  "Steer",
	//	Count: 1,
	//})
	//executeTemplate(templateText, Part{
	//	Name:  "Arrow",
	//	Count: 99999,
	//})

	type Subscriber struct {
		Name   string
		Active bool
		Rate   float64
	}

	templateText := "Name: {{.Name}}\n{{if .Active}}Rate: {{.Rate}}\n{{end}}"
	data := Subscriber{
		Name:   "WYM",
		Active: false,
		Rate:   4.5,
	}
	executeTemplate(templateText, data)

	data = Subscriber{
		Name:   "GUU",
		Active: true,
		Rate:   4.99,
	}
	executeTemplate(templateText, data)
}
