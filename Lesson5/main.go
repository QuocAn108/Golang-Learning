package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
	sum = 1
	for sum < 1000 { // while loop, do not using while in Go
		sum += sum
	}
	fmt.Println(sum)
	for { // infinite loop <=> while(true)

	}

}
