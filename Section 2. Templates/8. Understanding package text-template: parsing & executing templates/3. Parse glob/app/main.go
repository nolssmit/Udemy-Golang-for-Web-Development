// https://pkg.go.dev/text/template#pkg-index
package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	// Parse all templates in the "templates" directory
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Execute a spesific template
	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}

}

// Once you parese files you get a value of a pointer to a template.
// func ParseFiles(filenames ...string) (*Template, error)

// Then you have access to *Template's functions, e.g.
// func (t *Template) Execute(wr io.Writer, data any) error
