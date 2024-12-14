package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length
	s = s[:0]
	printSlice(s)

	// Extend its length
	s = s[:4]
	printSlice(s)

	// Drop its first two values
	s = s[2:]
	printSlice(s)
	var e []int
	fmt.Println(e, len(e), cap(e))
	if e == nil {
		fmt.Println("nil!")
	}

	a := make([]int, 5)
	printSlice2("a", a)

	b := make([]int, 0, 5)
	printSlice2("b", b)

	c := b[:2]
	printSlice2("c", c)

	d := c[2:5]
	printSlice2("d", d)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
func printSlice2(s string, x []int) {
	fmt.Println("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

//len is the number of elements in the slice
//cap is the number of elements in the underlying array, counting from the first element in the slice
//nil <=> null in other languages (null pointer)

//make([]T, len, cap) creates a slice of type T with a length of len and a capacity of cap
//s[a:b] creates a slice from a to b-1
//s[:b] creates a slice from the beginning to b-1
//s[a:] creates a slice from a to the end
//s[:] creates a slice from the beginning to the end
//nil slice has a length and capacity of 0
//nil slice is the zero value for a slice
