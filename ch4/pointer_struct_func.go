package ch4

import (
	"fmt"
	"strings"
)

func c() {
	i, j := 42, 2701

	p := &i         // 指向 i
	fmt.Println(*p) // 通过指针读取 i 的值
	*p = 21         // 通过指针设置 i 的值
	fmt.Println(i)  // 查看 i 的值

	p = &j         // 指向 j
	*p = *p / 37   // 通过指针对 j 进行除法运算
	fmt.Println(j) // 查看 j 的值
}

// 一个结构体（struct）就是一组字段（field）
type User struct {
	Age  int
	Name string
	Pwd  string
}

func G() {

	//数组
	var nameArrays [5]string

	nameArrays[0] = "test"
	nameArrays[1] = "test1"
	nameArrays[2] = "test2"
	nameArrays[3] = "test3"
	nameArrays[4] = "test4"

	values := [5]int{1, 2, 3, 2}
	fmt.Println(values)

	value := [...]int{123, 13, 213, 2131}
	fmt.Println(value)

	//切片
	//切片就像数组的引用
	//切片并不存储任何数据，它只是描述了底层数组中的一段。
	//更改切片的元素会修改其底层数组中对应的元素。
	//与它共享底层数组的切片都会观测到这些修改。
	// nil 切片的长度和容量为 0 且没有底层数组。
	//1.
	strings := nameArrays[1:4]
	strings[2] = "7"
	fmt.Println(strings)
	//2.888888809o00ol=[
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(len(q), cap(q), len(strings), cap(strings))
	//3
	ints := make([]int, 0, 5)
	fmt.Println(len(ints), cap(ints))

	for _, v := range q {
		fmt.Printf("2**%s = %d\n", "index", v)
	}
	// tips:Go 切片扩容策略：如果切片的容量小1024个元素，于是扩容的时候就翻倍增加容量。
	//上面那个例子也验证了这一情况，总容量从原来的4个翻倍到现在的8个。
	//一旦元素个数超过1024个元素，那么增长因子就变成 1.25，即每次增加原来容量的1/4。
	//本质引用数据  但是如果扩容了就不一样
	array := [4]int{10, 20, 30, 40}
	slice := array[0:2]           // 10 20
	newSlice := append(slice, 50) // 10 20 50
	newSlice[1] += 10             // 10 30 50
	// 这里slice=[10 30]，array=[10 30 50 40]，入坑！！！
}

func Slice_() {

	// 创建一个井字板（经典游戏）
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	//len3 cap 3
	// 两个玩家轮流打上 X 和 O
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		i2 := board[i]
		fmt.Printf("%s\n", strings.Join(i2, " "))
	}

}

var m map[string]int

func MapFunc() {
	m = make(map[string]int)
	m["lnh"] = 23
	var m2 = map[string]int{
		"rain": 123,
	}
	i, flag := m2["rain"]
	println(m, m2, i, flag)

	if _, flag := m2["rain"]; flag {
		fmt.Println("get a tips i, \n" +
			" if _, flag := m2[\"rain\"]; flag {}")
	}
}

func WordCount(s string) map[string]int {
	m = make(map[string]int)
	fields := strings.Fields(s)
	for _, field := range fields {
		value, flag := m[field]
		if flag {
			m[field] = value + 1
		} else {
			m[field] = 1
		}
	}
	return m
}

// 函数也是值。它们可以像其它值一样传递。
// 函数值可以用作函数的参数或返回值。
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// 函数闭包
func Adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// Fibonacci 函数返回一个闭包，该闭包返回一个斐波那契数列
func Fibonacci() func() int {
	a, b := 0, 1

	return func() int {
		result := a
		a, b = b, a+b
		return result
	}
}
