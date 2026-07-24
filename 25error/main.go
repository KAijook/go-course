package main

import (
	"errors"
	"fmt"
)

func f(arg int) error {
	if arg == 42 {

		return errors.New("wrong number")
	}

	return nil
}

var ErrOutOfTea = errors.New("no more tea available")

func makeTea(arg int) error {
	switch arg {
	case 1:
		return fmt.Errorf("making tea: %w", ErrOutOfTea)
	case 2:
		return ErrOutOfTea
	default:
		return nil
	}

}

func main() {

	i := 1
	fmt.Println(f(42))
	err := makeTea(i)
	if errors.Is(err, ErrOutOfTea) {
		fmt.Println("We should buy new tea!")
	}

}
