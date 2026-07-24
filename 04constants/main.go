package main

import "fmt"

const Pi = 3.14

func main() {
	v := 3.142
	fmt.Printf("v is of type %\n", v)

	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
