package main

import (
	"fmt"
	"sync"
)

type Ticket struct {
	mu     sync.Mutex
	amount int
}

func (t *Ticket) Buy(amount int, status chan bool) bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.amount >= amount {
		t.amount -= amount
		return true
	} else {
		fmt.Println("Hết vé")
		status <- false
		return false
	}
}

func main() {
	worker := 3
	customerAmount := make(chan int, 50)
	tickets := 10
	status := make(chan bool, 1)
	wg := sync.WaitGroup{}

	for i := range 50 {
		customerAmount <- i
	}
	close(customerAmount)

	ticket := &Ticket{
		amount: tickets,
	}

	for i := 1; i <= worker; i++ {

		wg.Go(func() {

			for customer := range customerAmount {

				if ticket.Buy(1, status) {
					fmt.Printf("worker %d đặt vé %d\n", i, customer)
				} else {

					break
				}

			}
		})
	}

	wg.Wait()
}
