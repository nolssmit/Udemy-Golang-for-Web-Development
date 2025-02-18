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
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", `Release self-focus; embrace other-focus.`)
	if err != nil {
		log.Fatalln(err)
	}
}

// In the `tpl.gohtml` file, there is a variable $wisdom that needs to be replaced with
//   the value of the string `Release self-focus; embrace other-focus.`.
// The template should replace the variable $wisdom with the provided string and print the result to the standard output.
// The main program will execute the parsed template from the file `tpl.gohtml` with the provided data.
