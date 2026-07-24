package main

import "fmt"

type bird struct {
	name string
	age  int
}

func (b bird) fly() {
	fmt.Println(b.name, "is flying")
}
func (b bird) sing() {
	fmt.Println(b.name, "is singing")
}

type dog struct {
	name string
	age  int
}

func main() {
	bird1 := bird{name: "Sparrow", age: 2}
	bird1.fly()
	bird1.sing()

	dog1 := dog{name: "Buddy", age: 3}
	fmt.Println(dog1.name, "is barking")

}
