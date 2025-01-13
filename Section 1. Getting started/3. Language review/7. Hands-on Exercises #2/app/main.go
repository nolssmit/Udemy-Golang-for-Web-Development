package main

import (
	"fmt"
)

// Syntax of functions
// func(receiver) indentifier(parameters) (returns) {code}

type person struct {
	fname string
	lname string
}

func (p person) pSpeak() string {
	return "Hello, my name is " + p.fname + " " + p.lname
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) pSpeak() string {
	return "Hello, my name is " + s.fname + " " + s.lname
}

func (s secretAgent) saSpeak() string {
	return "Stirred, not shaken"
}

func main() {
	p := person{fname: "Nols", lname: "Smit"}
	fmt.Println("p is of type", fmt.Sprintf("%T", p))
	fmt.Println("p has fields", p.fname, p.lname)
	fmt.Printf("Person introduced himself: %v\n", p.pSpeak())
	fmt.Println("----------------------")

	sa := secretAgent{person: person{fname: "Todd", lname: "McLeod"}, ltk: true}
	fmt.Println("sa is of type", fmt.Sprintf("%T", sa))
	fmt.Println("sa has fields", sa.fname, sa.lname, sa.ltk)
	fmt.Printf("Secret Agent introduced himself: %v\n", sa.pSpeak())
	fmt.Println("Secret agent say: ", sa.saSpeak())
}
/*
	create a struct that holds person fields
	create a struct that holds secret agent fields and embeds person type
	attach a method to person: pSpeak
	attach a method to secret agent: saSpeak
	Todd McLeod - Learn To Code Golang - Hands-On Exercises #1 Page 2
	create a variable of type person
	create a variable of type secret agent
	print a field from person
	run pSpeak attached to the variable of type person
	print a field from secret agent
	run saSpeak attached to the variable of type secret agent
	run pSpeak attached to the variable of type secret agent
*/