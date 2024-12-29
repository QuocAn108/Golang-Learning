package main

import (
	"fmt"
	"time"
)

func main() {
	n := 10000
	maxWorker := 5
	queueCh := make(chan int, n)
	for i := 1; i <= n; i++ {
		queueCh <- i
	}
	close(queueCh)
	for i := 0; i <= maxWorker; i++ {
		go func(count int) {
			time.Sleep(time.Second)
			fmt.Printf("Worker %d is crawling web url %d\n", count, <-queueCh)
		}(i)
	}
	time.Sleep(time.Second * 100)
}
