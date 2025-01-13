package main

import (
	"fmt"
)

type person struct {
	fname string
	lname string
}

func (p person) pSpeak() string {
	return fmt.Sprintf("I am %s %s", p.fname, p.lname)
}
//------------------------------------------------------------

type secretAgent struct {
	person
	licenseToKill bool
}

func (s secretAgent) pSpeak() string {
	return fmt.Sprintf("I am %s %s and have license to kill: %t", s.fname, s.lname, s.licenseToKill)
}
//------------------------------------------------------------

type human interface {
	speak() string
}

func saySomething(h human) {
	h.speak()
}
//------------------------------------------------------------

func main() {

	p := person{fname: "Nols", lname: "Smit"}
	fmt.Println("p is of type", fmt.Sprintf("%T", p))
	fmt.Println("p has fields", p.fname, p.lname)
	fmt.Printf("Person introduced himself: %v\n", p.pSpeak())
	fmt.Println("-------------------------------------------------------")

	sa := secretAgent{person: person{fname: "Todd", lname: "McLeod"}, licenseToKill: true}
	fmt.Println("sa is of type", fmt.Sprintf("%T", sa))
	fmt.Println("sa has fields", sa.fname, sa.lname, sa.licenseToKill)
	fmt.Printf("Secret Agent introduced himself: %v\n", sa.pSpeak())
	fmt.Println("Secret agent say: ", sa.pSpeak())
	fmt.Println("-------------------------------------------------------")	
	human1 := secretAgent{person: person{fname: "James", lname: "Bond"}, licenseToKill: true}
	fmt.Println("Human1 say: ", human1.pSpeak())
	fmt.Println("-------------------------------------------------------")
	human2 := person{fname: "Miss", lname: "Moneypenny"}
	fmt.Println("Human2 say: ", human2.pSpeak())
	fmt.Println("-------------------------------------------------------")
}

/*
	create an interface type that both person and secretAgent implement
	declare a func with a parameter of the interface’s type
	call that func in main and pass in a value of type person
	call that func in main and pass in a value of type secretAgent
*/
