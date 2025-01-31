package main

import "fmt"

func main() {
	defer fmt.Println("World") // defer will execute after the main function is done
	fmt.Println("Hello")

	fmt.Println("Counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i) // defer will execute in reverse order
	}
	fmt.Println("Done")
}
