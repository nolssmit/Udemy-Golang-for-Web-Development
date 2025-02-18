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
	sages := map[string]string{
		"Ghandi":   "The Sage of India",
		"MLK":      "The Father of Nonviolence",
		"Buddha":   "The Buddha of Compassion",
		"Jesus":    "The Christ of Love",
		"Muhammed": "The Prophet of Islam",
	}

	err := tpl.Execute(os.Stdout, sages)

	if err != nil {
		log.Fatalln(err)
	}
}
