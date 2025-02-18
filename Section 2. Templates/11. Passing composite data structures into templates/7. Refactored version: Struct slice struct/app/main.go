// https://pkg.go.dev/text/template#pkg-index
package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type items struct {
	Wisdom    []sage
	Transport []car
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

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

	f := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	c := car{
		Manufacturer: "Honda",
		Model:        "Civic",
		Doors:        4,
	}

	sages := []sage{b, g, m}
	cars := []car{f, c}

	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		sages,
		cars,
	}

	err := tpl.Execute(os.Stdout, data)

	if err != nil {
		log.Fatalln(err)
	}
}
