package main

import "fmt"

func main() {

	s := map[string]int{
		"first":  1,
		"second": 2,
	}
	for key, value := range s {
		fmt.Println(key, value)
	}
	for key := range s {
		fmt.Println(s[key])
	}

	array := [3]int{1, 2, 3}
	for index, value := range array {
		fmt.Println(index, value)
	}

	slice := []int{4, 5, 6}
	for _, value := range slice {
		fmt.Println(value)
	}

	for a, b := range "go" {
		fmt.Println(a, b)
	}
}
