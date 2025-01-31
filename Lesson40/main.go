package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	n := 10000
	maxWorker := 5
	queueCh := make(chan int, n)
	wg := new(sync.WaitGroup)
	for i := 1; i <= n; i++ {
		queueCh <- i
	}
	close(queueCh)
	for i := 0; i <= maxWorker; i++ {
		wg.Add(1)
		go func(count int) {
			for v := range queueCh {
				time.Sleep(time.Millisecond * 20)
				fmt.Printf("Worker %d is crawling web url %d\n", count, v)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}

//using waitgroup to wait for all goroutines to finish before exiting the main function
//using channel to pass values to goroutines
//using close to close the channel
