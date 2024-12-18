package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 { // method
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func Abs(v Vertex) float64 { // function
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) { // function
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Scale(f float64) { // method
	v.X = v.X * f
	v.Y = v.Y * f
}
func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
