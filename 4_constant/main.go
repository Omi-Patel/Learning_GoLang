package main

import "fmt"

const age = 22 // constant can be declare outside the function

// age := 22 // this syntax is not working outside the function in case of constant

func main() {
	const name = "go lang"

	// name = "javascript" // you can not assign a value to constant variable

	fmt.Println(age)

	// constant grouping
	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)
}
