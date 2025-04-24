// https://medium.com/@souravchoudhary0306/why-golang-favour-composition-over-inheritance-6196342d7cfe
// https://medium.com/deep-golang/why-golang-choose-composition-as-its-base-rather-than-inheritance-1225d22a4798

package main

import (
	"log"
	"os"
	"text/template"
)

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	y := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				{Number: "101", Name: "Introduction to Computer Science", Units: "4"},
				{Number: "102", Name: "Data Structures", Units: "3"},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				{Number: "201", Name: "Algorithms", Units: "3"},
				{Number: "202", Name: "Operating Systems", Units: "3"},
			},
		},
		Summer: semester{
			Term: "Summer",
			Courses: []course{
				{Number: "301", Name: "Database Management", Units: "3"},
				{Number: "302", Name: "Computer Networks", Units: "3"},
			},
		},
	}

	err := tpl.Execute(os.Stdout, y)
	if err != nil {
		log.Fatal(err)
	}
}
