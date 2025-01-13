package main

import (
	"fmt"
)

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

// Syntax of functions
// func(receiver) indentifier(parameters) (returns) {code}
// Attach a function of certain type to a struct variable
func (p person) speak() {
	fmt.Println(p.fname, p.lname, `says, "Good moring, James."`)
}

func(sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `says, "Shaken, not stirred."`)
}
func main() {
	p1 := person{
		"Penny", 
		"Monneypenny",
	}
	p1.speak()

	secretAgent1 := secretAgent{
		person {
			"James",
			"Bond",
		},
		true,
	}
	secretAgent1.speak()
	// Drill down to the inner type
	secretAgent1.person.speak()
}

