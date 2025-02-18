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
	g1 := struct {
		Score1 int
		Score2 int
	}{
		7,
		9,
	}

	err := tpl.Execute(os.Stdout, g1)
	if err != nil {
		log.Fatal(err)
	}
}
