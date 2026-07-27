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
				select {
				case orderQueue <- i:
					fmt.Println("Gửi đơn thành công")
					sent = true
				default:
					fmt.Println("Đang đầy hàng chờ")
					time.Sleep(100 * time.Millisecond)
				}
			}
		}
		close(orderQueue)
	})

	for i := 1; i <= worker; i++ {
		wg.Go(func() {

			for order := range orderQueue {
				time.Sleep(300 * time.Millisecond)
				<-orderQueue
				fmt.Printf("Đầu bếp đã hoàn thành đơn %d\n", order)
			}
		})

	}

	wg.Wait()
}
