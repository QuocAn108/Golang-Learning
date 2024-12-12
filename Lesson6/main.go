package main

import (
	"fmt"
	"math"
)

func srqt(x float64) string {
	if x < 0 { //condition do not have to be in parentheses
		return srqt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim { //v is a variable that is only available in the if scope
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}
func main() {
	fmt.Println(srqt(2), srqt(-4))
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
}
