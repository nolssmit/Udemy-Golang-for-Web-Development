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

type user struct {
	Name  string
	Motto string
	Admin bool
}

func main() {
	U1 := user{Name: "Aditya", Motto: "Always be yourself", Admin: true}
	u2 := user{Name: "John", Motto: "Be the best you can be", Admin: false}
	u3 := user{Name: "", Motto: "Love what you do", Admin: false}

	users := []user{U1, u2, u3}

	err := tpl.Execute(os.Stdout, users)
	if err != nil {
		log.Fatal(err)
	}
}
