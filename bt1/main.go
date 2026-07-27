package main

import (
	"fmt"
	"sync"
)

type Wallet struct {
	mu      sync.Mutex
	balance int
}

func (w *Wallet) Deposit(amount int, transaction chan int) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.balance += amount
	transaction <- w.balance
}

func (w *Wallet) Withdraw(amount int, transaction chan int) {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.balance >= amount {
		w.balance -= amount
		transaction <- w.balance
	} else {
		fmt.Println("Số dư không đủ")
	}
}

func main() {
	wg := sync.WaitGroup{}

	transaction := make(chan int, 1)

	wallet := &Wallet{
		balance: 1000,
	}
	DepositWorker := 5
	WithdrawWorker := 3

	for i := 0; i < WithdrawWorker; i++ {
		wg.Go(func() {

			wallet.Deposit(500000, transaction)
			fmt.Printf("Current Balance: %d\n", <-transaction)
		})

	}

	for i := 0; i < DepositWorker; i++ {
		wg.Go(func() {
			wallet.Withdraw(200000, transaction)
			fmt.Printf("Current Balance: %d\n", <-transaction)
		})
	}

	wg.Wait()
	close(transaction)

}
