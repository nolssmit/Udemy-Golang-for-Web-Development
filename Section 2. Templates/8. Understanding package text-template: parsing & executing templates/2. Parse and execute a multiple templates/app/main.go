// https://pkg.go.dev/text/template#pkg-index
package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("one.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal(err)
	}
	//                ------------------------
	tpl, err = template.ParseFiles("two.gohtml", "vespa.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	// Specify which template to execute
	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.Execute(os.Stdout, "two.gohtml")
	if err != nil {
		log.Fatal(err)
	}

}

// Once you parese files you get a value of a pointer to a template.
// func ParseFiles(filenames ...string) (*Template, error)

// Then you have access to *Template's functions, e.g.
// func (t *Template) Execute(wr io.Writer, data any) error
