package main

import "fmt"

func fusion(x, y string) string {
	return x + y
}
func triple(x, y, z int) int {
	return x + y + z
}

func main() {
	fmt.Println(fusion("go", "ku"))
	fmt.Println(triple(1, 2, 3))
}
