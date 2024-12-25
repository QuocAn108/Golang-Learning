package main

import "fmt"

// index returns the index of the first instance of e in s, or -1 if e is not present in s.
func index[T comparable](s []T, e T) int {
	for i, v := range s {
		if v == e {
			return i
		}
	}
	return -1
}

func main() {
	si := []int{10, 20, 15, -10}
	fmt.Println(index(si, 15)) // 2

	ss := []string{"cat", "dog", "bird"}
	fmt.Println(index(ss, "cc")) // -1
}

//type parameters are declared in square brackets after the function name
//comparable is a type constraint that specifies that the type parameter T must be a comparable type
//The index function takes a slice s of type T and a value e of type T and returns an int.
//The function iterates over the elements of the slice s using a for loop with range.
