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

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	gandhi := sage{
		Name:  "Gandhi",
		Motto: "Be the change you want to see",
	}

	mlk := sage{
		Name:  "Martin Luther King Jr.",
		Motto: "Do not stand idly by, let us take our stand",
	}

	jesus := sage{
		Name:  "Jesus Christ",
		Motto: "Love your enemies, do good to those who hate you",
	}

	sages := []sage{buddha, gandhi, mlk, jesus}

	err := tpl.Execute(os.Stdout, sages)

	if err != nil {
		log.Fatalln(err)
	}
}
