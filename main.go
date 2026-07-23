package main

import "fmt"

var i int = 42
var f float64 = float64(i)
var z uint = uint(f)

func main() {
	fmt.Println(i, f, z)
}
