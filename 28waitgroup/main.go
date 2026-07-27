package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		time.Sleep(time.Second)
		fmt.Println("worker", id, "processing job", j)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	var wg sync.WaitGroup

	for i := range 10 {
		jobs <- i
	}
	close(jobs)

	for i := 0; i < 3; i++ {
		wg.Go(func() {
			worker(i, jobs, results)
		})
	}
	wg.Wait()
	close(results)
	for result := range results {
		fmt.Println("result:", result)
	}
}
