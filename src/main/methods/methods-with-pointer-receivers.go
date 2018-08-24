package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// * 如果去掉的话 值将不会被改变。
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v, v.Abs())
}
