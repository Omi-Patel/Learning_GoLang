package main

import "fmt"

func main() {
	// age := 6

	// if age >= 18 {
	// 	fmt.Println("person is an adult!!")
	// } else {
	// 	fmt.Println("person is not an adult!!")
	// }

	// if age >= 18 {
	// 	fmt.Println("person is an adult!!")
	// } else if age >= 12 {
	// 	fmt.Println("person is a teenager!!")
	// } else {
	// 	fmt.Println("person is a kid!")
	// }

	// declaring and initializing variable inside if statement
	if age := 15; age >= 18 {
		fmt.Println("person is an adult!!", age)
	} else if age >= 12 {
		fmt.Println("person is a teenager!", age)
	}
}

// go does not have ternary operation!!
