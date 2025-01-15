package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := os.Args[1]
	fmt.Println(os.Args[0]) // program you executed
	fmt.Println(os.Args[1]) // additional data
	str := fmt.Sprintf("%s", `

	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Hello World!</title>
	</head>
	<body>
		<h1>Hello `+name+`</h1>
	</body>
	</html>
	`)

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(str))
}
// Test with: go run main.go "Nols Smit"  > index.html 	