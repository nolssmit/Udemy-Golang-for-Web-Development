// https://pkg.go.dev/text/template#pkg-index
package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}
}
// In the `tpl.gohtml` file, there is a placeholder `<h1>{{.}}</h1>` that needs to be replaced with 
//   the value of the variable `42`. 
// The `main` function should execute the template and print the result to the standard output.