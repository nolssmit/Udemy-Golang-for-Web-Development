// https://pkg.go.dev/text/template#pkg-index
package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

// create a FuncMap to register functions.
// "uc" is what the func will be called in the template.
// "uc" is the ToUpper func from package strings
// "ft" is the func I declared.
// "ft" slices a string, returning first three characters.
type sage struct {
	Name  string
	Motto string
}

type items struct {
	Persons []sage
}

func firstFive(s string) string {
	s = strings.TrimSpace(s)
	s = s[:6]
	return s
}

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ff": firstFive,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

var tpl *template.Template

func main() {
	b := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	g := sage{
		Name:  "Gandhi",
		Motto: "Be the change you want to see",
	}

	m := sage{
		Name:  "Martin Luther King Jr.",
		Motto: "Do not stand idly by, let us take our stand",
	}

	sages := []sage{b, g, m}

	data := items{
		Persons: sages,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)

	if err != nil {
		log.Fatalln(err)
	}
}
