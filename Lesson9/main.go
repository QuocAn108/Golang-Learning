package main

import "fmt"

func main() {
	i, j := 42, 2701
	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i
	p = &j          // point to j
	*p = *p / 37    // divide j through the pointer
	fmt.Println(j)  // see the new value of j
}

//unlike C, Go has no pointer arithmetic.
//example: p++ is illegal in Go.
//Go has garbage collection, so it is safe to return a pointer to local variable.
