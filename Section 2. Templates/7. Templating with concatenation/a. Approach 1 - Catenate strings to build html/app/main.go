package main

import (
	"fmt"
)

func main() {
	name := "John Doe"

	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Hello World!</title>
	</head>
	<body>
		<h1>Hello ` + name + `</h1>
	</body>
	</html>
	`

	fmt.Println(tpl)
}
// go run main.go > index.html