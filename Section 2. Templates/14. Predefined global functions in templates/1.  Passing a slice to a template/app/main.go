// https://pkg.go.dev/github.com/aditya43/golang-core/079-Predefined-Global-Functions
// https://godoc.org/text/template#hdr-Functions
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
	xs := []string{"zero", "one", "two", "three", "four", "five"}
	err := tpl.Execute(os.Stdout, xs)
	if err != nil {
		log.Fatal(err)
	}
}
