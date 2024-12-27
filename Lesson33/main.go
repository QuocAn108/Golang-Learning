package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("Hello")
}

//Goroutines in go is a lightweight thread of execution. It is a function that runs concurrently with other functions. To create a goroutine, we use the keyword go followed by the function invocation. In the above example, we have two functions say and main. The main function is the entry point of the program. The say function prints the string passed to it five times with a delay of 100 milliseconds. In the main function, we have two function invocations. The first function invocation is a goroutine that runs concurrently with the main function. The second function invocation is a normal function call. When we run the program, the output will be as follows:
