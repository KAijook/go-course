package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	orderQueue := make(chan int, 5)
	worker := 3
	wg := sync.WaitGroup{}

	wg.Go(func() {
		for i := range 30 {
			sent := false
			for !sent {
				if len(orderQueue) < cap(orderQueue) {
					orderQueue <- i
					fmt.Println("Đơn hàng đã được gửi")
					sent = true
				} else {
					time.Sleep(300 * time.Millisecond)
					fmt.Println("Đang đầy hàng chờ")
				}
			}
		}
		close(orderQueue)
	})

	for i := 1; i <= worker; i++ {
		wg.Go(func() {

			for order := range orderQueue {
				time.Sleep(300 * time.Millisecond)

				fmt.Printf("Đầu bếp đã hoàn thành đơn %d\n", order)
			}
		})

	}

	wg.Wait()
}
