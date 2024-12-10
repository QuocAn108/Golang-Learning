package main

import "fmt"

func swap(x,y string) (string, string) { //declare for multiple return values
	return y, x
}

func main() {
	a, b := swap("hello", "world"); //using function with multiple return values
	fmt.Println(a, b)
}
//go don't use ';' to end the statement
//exported names start with a capital letter