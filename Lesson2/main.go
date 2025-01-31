package main

import "fmt"

func split(sum int) (x, y int) { // x, y are named return values
	x = sum * 4 / 9
	y = sum - x
	return // return statement without arguments returns the named return values
}

func main() {
	fmt.Println(split(17))
}
