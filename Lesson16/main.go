package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 { //fn is name of variable, func(float64, float64) float64 is type of variable
	return fn(3, 4)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))      // 13
	fmt.Println(compute(hypot))    // 5
	fmt.Println(compute(math.Pow)) // 81
}

//function closures is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
//For example, the adder function returns a closure. Each closure is bound to its own sum variable.
