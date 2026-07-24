package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp:", a)
	a[2] = 1
	fmt.Println("set:", a)
	fmt.Println("get:", a[2])
	fmt.Println("len:", len(a))
	a = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", a)
	a = [5]int{1: 100}
	fmt.Println("dcl:", a)
	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("2d: ", twoD)
}
