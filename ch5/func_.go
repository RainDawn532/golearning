package ch5

import (
	"fmt"
	"reflect"
)

type Person struct {
}
type Model struct {
	Age  int
	Name string
	Person
}

func (v Model) GetName() string {
	str := fmt.Sprintf("%d", v.Age)
	return v.Name + str
}
func (v *Model) GetPointerName() string {
	str := fmt.Sprintf("%d", v.Age)
	return v.Name + str
}
func (v Person) GetPersonName() {

}
func (v *Person) GetPointerPersonName() {

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

func MethodSet(a interface{}) {
	t := reflect.TypeOf(a)
	for i, n := 0, t.NumMethod(); i < n; i++ {
		m := t.Method(i)
		fmt.Println(m.Name, m.Type)
	}
}
