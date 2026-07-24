package main

import (
	"cmp"
	"fmt"
	"slices"
)

func lenCompare(s1, s2 string) int {
	return cmp.Compare(len(s1), len(s2))
}

func main() {
	s := []int{1, 3, 2, 5, 4}
	slices.Sort(s)
	fmt.Println(s)

	b := []string{"go", "is", "awesome"}
	slices.Sort(b)
	fmt.Println(b)
	fmt.Println(slices.IsSorted(s))

	words := []string{"go", "is", "awesome"}
	slices.SortFunc(words, lenCompare)
	fmt.Println(words)
}
