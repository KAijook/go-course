package main

import "fmt"

func multiple(factor int) func(int) int {
	return func(x int) int {
		return factor * x
	}
}

func main() {

	double := multiple(2)

	fmt.Println(double(3))

	newInts := multiple(3)
	fmt.Println(newInts(4))
}
