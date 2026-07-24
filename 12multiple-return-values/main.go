package main

import "fmt"

func vals() (int, int, string) {
	return 3, 7, "hello"
}

func main() {

	a, b, c := vals()
	fmt.Println(a)
	fmt.Println(c)
	fmt.Println(b)

	_, _, c = vals()
	fmt.Println(c)
}
