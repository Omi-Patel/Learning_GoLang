package main

import (
	"fmt"
	"maps"
)

func main() {

	//            [key]value
	m := make(map[string]string)

	// setting an element
	m["name"] = "golang"
	m["area"] = "backend"

	// getting an element
	fmt.Println(m["name"], m["area"])

	// IMP : if key doesn't exist in map then it returns zero value

	fmt.Println(len(m))

	delete(m, "area")

	fmt.Println("after delete ", m)

	clear(m)
	fmt.Println("after clear ", m)

	// creating map without make()
	mp := map[string]int{"price": 40, "phones": 3}
	// fmt.Println(mp)

	v, ok := mp["price"]

	fmt.Println(v) // v -> returns the value of key if it present in map
	// ok -> return true if key is present in map otherwise returns false

	if ok {
		fmt.Println("all ok")
	} else {
		fmt.Println("not ok")
	}

	// compare two maps

	m1 := map[string]int{"age": 22}
	m2 := map[string]int{"age": 22}
	fmt.Println(maps.Equal(m1, m2))
}
