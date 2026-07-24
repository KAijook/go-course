package main

import (
	"fmt"
	"slices"
)

func main() {
	s := []int{1, 3, 2, 5, 4}
	slices.Sort(s)
	fmt.Println(s)

	b := []string{"go", "is", "awesome"}
	slices.Sort(b)
	fmt.Println(b)
	fmt.Println(slices.IsSorted(s))
}
