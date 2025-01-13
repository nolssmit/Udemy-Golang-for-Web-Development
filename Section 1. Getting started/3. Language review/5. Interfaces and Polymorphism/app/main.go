package main

import "fmt"

//	"fmt"

// Syntax of functions
// func(receiver) indentifier(parameters) (returns) {code}
// Attach a function of certain type to a struct variable

type person struct {
	fname string
	lname string
}

func (p person) speak()  {
	//fmt.Println(p.fname, p.lname, `says, "Good moring, James."`)
	fmt.Println(p.fname + " " + p.lname + "says, Good moring, James.")
}

//-------------------------------------------------------------------

type secretAgent struct {
	person
	licenseToKill bool
}

func (sa secretAgent) speak() {
	//fmt.Println(sa.fname, sa.lname, `says, "Shaken, not stirred."`)
	fmt.Println(sa.fname + " " + sa.lname + "says, Shaken, not stirred.")
}

//-------------------------------------------------------------------

// A interface can implement multiple types
type human interface {
	speak()
}

func saySomething(h human) string {
	return fmt.Sprintf("%v",h.speak())
}

//-------------------------------------------------------------------

func main() {
	p1 := person{
		"Penny",
		"Monneypenny",
	}

	secretAgent1 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}

	p1.speak()
	secretAgent1.speak()
	fmt.Println("---------------------------------")
	// PolyMorphism
	saySomething(p1)
	saySomething(secretAgent1)
}
