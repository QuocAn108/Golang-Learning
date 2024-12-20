package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// f = i.(float64) // panic
	// fmt.Println(f)
}

//type assertion provides access to an interface value's underlying concrete value.
// t := i.(T) asserts that the interface value i holds a concrete type T and assigns the underlying T value to the variable t.
// If i does not hold a T, the statement will trigger a panic.
// To test whether an interface value holds a specific type, a type assertion can return two values: the underlying value and a boolean value that reports whether the assertion succeeded.
