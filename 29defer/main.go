package main

import (
	"fmt"
)

func safeDivider(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("division by zero")
		}
	}()
	return a / b, nil
}

func main() {
	result, err := safeDivider(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
