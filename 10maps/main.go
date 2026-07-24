package main

import (
	"fmt"
	"maps"
)

func main() {
	a := make(map[string]int)
	a["one"] = 1
	a["two"] = 2
	fmt.Println(a)

	fmt.Println(a["one"])

	delete(a, "one")
	fmt.Println(a)

	a["three"] = 3
	clear(a)
	fmt.Println(a)

	b := map[string]int{
		"four": 4,
		"five": 5,
	}
	fmt.Println(b)

	c := map[string]int{
		"four": 4,
		"five": 5,
	}
	if maps.Equal(b, c) {
		fmt.Println("b and c are equal")
	}
}
