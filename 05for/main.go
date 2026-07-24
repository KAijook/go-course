package main

import "fmt"

func main() {
	c := 3
	for c <= 5 {
		fmt.Print(c)
		c++
	}

	for i := 0; i < 5; i++ {
		fmt.Print(i)
	}
	for i := range 3 {
		fmt.Print(i)
	}
}
