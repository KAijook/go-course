package main

import "fmt"

type Dog struct {
}

func (d Dog) Bark() {
	fmt.Println("Woof!")
}
func (d Dog) Sound() {
	d.Bark()
}

type Cat struct{}

func (c Cat) Meow() {
	fmt.Println("Meow!")
}
func (c Cat) Sound() {
	c.Meow()
}

type Animal interface {
	Sound()
}

func isDog(a Animal) bool {
	_, ok := a.(Dog)
	return ok
}

func main() {
	dog := Dog{}
	cat := Cat{}
	dog.Sound()
	cat.Sound()
	if isDog(dog) {
		fmt.Println("It's a dog!")
	}
	if isDog(cat) {
		fmt.Println("It's a dog!")
	} else {
		fmt.Println("It's not a dog!")
	}
}
