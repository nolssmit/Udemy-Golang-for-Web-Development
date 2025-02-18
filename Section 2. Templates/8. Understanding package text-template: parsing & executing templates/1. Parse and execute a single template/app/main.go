// https://pkg.go.dev/text/template#pkg-index
package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Once you parese files you get a value of a pointer to a template.
// func ParseFiles(filenames ...string) (*Template, error)

// Then you have access to *Template's functions, e.g.
// func (t *Template) Execute(wr io.Writer, data any) error
