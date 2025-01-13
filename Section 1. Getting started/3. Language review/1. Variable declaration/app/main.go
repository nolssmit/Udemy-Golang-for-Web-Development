package main

import (
	"fmt"
)

// Variable declaration with package level scope.
var y int

func main() {
	// Short variable declaration, with block leve scope. 
	// The declaration figures out the type. 
	// It can only be used in a code block, n	ot outside the functions
	x := 10
	fmt.Printf("The type of x is %T\n", x)
	fmt.Printf("The value of x is %v\n", x)

	// Printing out the default value of a package variable.
	fmt.Printf("The default value of y is %v\n", y) 
}