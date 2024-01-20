package ch5

import (
	"fmt"
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
	X, Y int
}

func Abs(v Vertex) int {
	return v.X
}

func Scale1(v Vertex) {
	v.X = v.X * v.X
	v.Y = v.Y * v.Y
}

func Scale(v *Vertex) {
	v.X = v.X * v.X
	v.Y = v.Y * v.Y
}
