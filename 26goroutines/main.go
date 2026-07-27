package main

import (
	"fmt"
)

func main() {
	c1 := 0
	sum := make(chan int)
	for i := 0; i < 100; i++ {

		go func() {

			sum <- i
		}()

	}
	c1 += <-sum
	fmt.Println(c1)

}
