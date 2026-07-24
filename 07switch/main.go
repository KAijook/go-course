package main

import "fmt"

func main() {
	status := "loading"

	switch status {
	case "loading":
		fmt.Println("The status is loading")
	case "ready":
		fmt.Println("The status is ready")
	case "error":
		fmt.Println("The status is error")
	case "complete":
		fmt.Println("The status is complete")
	default:
		fmt.Println("The status is unknown")
	}
}
