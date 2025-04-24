// https://medium.com/@souravchoudhary0306/why-golang-favour-composition-over-inheritance-6196342d7cfe
// https://medium.com/deep-golang/why-golang-choose-composition-as-its-base-rather-than-inheritance-1225d22a4798

package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

type doubleZero struct {
	person
	LicenseToKill bool
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	p1 := doubleZero{
		person{
			Name: "James Bond",
			Age:  30,
		},
		false,
	}

	err := tpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatal(err)
	}
}
