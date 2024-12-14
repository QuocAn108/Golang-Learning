package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] //slice of primes from index 1 to 4, remove element at first index
	fmt.Println(s)            //3 5 7

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names) //[John Paul George Ringo]
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)  //[John Paul] [Paul George]
	b[0] = "XXX"       //change the value of b[0] to XXX and it will reflect in a[1] as well "2 chang tro 1 nang in vnese"
	fmt.Println(a, b)  //[John XXX] [XXX George]
	fmt.Println(names) //[John XXX George Ringo]

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q) //[2 3 5 7 11 13]

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r) //[true false true true false true]

	d := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(d) //[{2 true} {3 false} {5 true} {7 true} {11 false} {13 true}]
}

//slice is array without a fixed length
//slice is a reference to a section of an array
//slice literal is like an array literal without the length
//slice defaults: low: 0, high: length of the array
