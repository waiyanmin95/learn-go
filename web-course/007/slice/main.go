package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
	//tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	f := []string{"apple", "grape", "mango"}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer nf.Close()

	err = tpl.Execute(nf, f)
	if err != nil {
		log.Fatalln(err)
	}
}

//
//package main
//
//import (
//	"html/template"
//	"os"
//)
//
//func main() {
//	// Define the data to be rendered in the template
//	data := []string{"apple", "banana", "cherry"}
//
//	// Define the template
//	tmpl, err := template.New("example").Parse(`
//		<ul>
//			{{range .}}
//				<li>{{.}}</li>
//			{{end}}
//		</ul>
//	`)
//	if err != nil {
//		panic(err)
//	}
//
//	// Render the template with the data
//	err = tmpl.Execute(os.Stdout, data)
//	if err != nil {
//		panic(err)
//	}
//}
