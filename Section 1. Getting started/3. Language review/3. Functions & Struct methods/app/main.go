package main

import (
	"fmt"
)

type person struct {
	fname string
	lname string
}

// Syntax of functions
// func(receiver) indentifier(parameters) (returns) {code}
// Attach a function of certain type to a struct variable
func (p person) speak() {
	fmt.Println(p.fname, p.lname, "says, Good moring, James.")
	
}
func main() {
	p1 := person{
		"Penny", 
		"Monneypenny",
	}
	p1.speak()
}

