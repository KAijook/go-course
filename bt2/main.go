package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Ticket struct {
	mu     sync.Mutex
	amount int
}

func (t *Ticket) Buy(amount int) bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.amount >= amount {
		t.amount -= amount
		return true
	} else {
		fmt.Println("Hết vé")

		time.Sleep(300 * time.Millisecond)
		return false
	}
}

func main() {
	worker := 3
	customerAmount := make(chan int, 50)
	tickets := 10
	var status atomic.Bool
	status.Store(true)

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

				if ticket.Buy(1) {
					fmt.Printf("worker %d đặt vé %d\n", i, customer)
				} else {

					break
				}

			}
		})
	}

	wg.Wait()
}
