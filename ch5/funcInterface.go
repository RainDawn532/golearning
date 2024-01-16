package ch5

import (
	"fmt"
	"math"
)

type Model struct {
	Age  int
	Name string
}

func (v Model) GetName() string {
	str := fmt.Sprintf("%d", v.Age)
	return v.Name + str
}
func GetName(v Model) string {
	str := fmt.Sprintf("%d", v.Age)
	return v.Name + str
}

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
