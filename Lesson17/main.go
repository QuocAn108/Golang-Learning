package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 { //method
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v Vertex) Abs() float64 { //method
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func vAbs(v Vertex) float64 { //function
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
	v := Vertex{3, 4}
	fmt.Println(vAbs(v))
	fmt.Println(v.Abs())
}

//methods is just a function with a receiver argument.
