package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

// at 2021-08-01 15:00:00.000000 +0000 UTC m=+0.000000000, it didn't work
//cannot use MyError literal (type *MyError) as type error in return argument
//can customize error message by implementing Error() method
