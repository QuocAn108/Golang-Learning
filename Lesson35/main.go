package main

import "fmt"

func main() {
	ch := make(chan int, 2) // buffered channel with capacity 2, don't need to use goroutine or func to receive the value from the channel
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
