package main

type List[T any] struct {
	next  *List[T]
	value T
}

func main() {}
