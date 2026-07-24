package main

import (
	"fmt"
)

func main() {
	var s []string
	fmt.Println(s, len(s), cap(s))

	s = make([]string, 3)
	fmt.Println(s, len(s), cap(s))

	s = append(s, "a")
	fmt.Println(s, len(s), cap(s))
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s, len(s), cap(s))
}
