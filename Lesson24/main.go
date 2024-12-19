package main

import "fmt"

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

//empty interface can hold values of any type, because every type implements at least zero methods.
//Empty interfaces are used by code that handles values of unknown type.
