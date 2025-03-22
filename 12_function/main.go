package main

import "fmt"

// func add(a, b int) int {
// 	return a + b
// }

// func getLanguages() (string, string, string, int) {
// 	return "java", "golang", "javascript", 22
// }

func processIt(fn func(a int) int) int {
	return fn(1)
}

func main() {
	// sum := add(1, 2)
	// fmt.Println(sum)

	// lang1, lang2, lang3, _ := getLanguages()
	// fmt.Println(lang1, lang2, lang3)

	fn := func(a int) int {
		return a
	}

	fmt.Println(processIt(fn))
}
