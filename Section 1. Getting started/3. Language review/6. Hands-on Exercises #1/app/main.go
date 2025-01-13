package main

import (
	"fmt"
)

// Syntax of functions
// func(receiver) indentifier(parameters) (returns) {code}

type square struct {
	side float64
}

func (s square) area() float64 {
	return s.side * s.side
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(s shape) float64 {
	return s.area()
}

func main() {
	sq := square{10}
	ci := circle{10}

	sqa := info(sq)
	cia := info(ci)

	fmt.Println("Area of square is: ", sqa)
	fmt.Println("Area of circle is: ", cia)
}
/*
create a type square
create a type circle
attach a method to each that calculates area and returns it
create a type shape which defines an interface as anything which has the area method
create a func info which takes type shape and then prints the area
create a value of type square
create a value of type circle
use func info to print the area of square
use func info to print the area of circle
*/