package main

import (
	"fmt"
	"golearning/ch5"
)

func main() {
	//ch1.SayHello()
	//swap, s := ch1.Swap(5, 8)
	//fmt.Println(swap, s)
	//defer 语句会将函数推迟到外层函数返回之后执行。
	//推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用
	//defer 栈
	//推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。
	//defer fmt.Println("before test")
	//defer fmt.Println(" after test")
	////i := ch2.DeferFunc()
	////fmt.Println(i)
	//panic("ecption")
	//ch3.F()
	//var user ch4.User
	//user.Name = "rainDawn"
	//user.Pwd = "12312"
	//user.Age = 123
	//fmt.Println(user)
	//ch4.G()
	//ch4.MapFunc()
	//wordMap := ch4.WordCount("foo bas as foo sau asau as foo")
	//fmt.Println(wordMap)
	////fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
	//
	//pos, neg := ch4.Adder(), ch4.Adder()
	//fmt.Print(pos(1), neg(2))

	//// 创建斐波那契数列的闭包
	//fib := ch4.Fibonacci()
	//
	//// 输出斐波那契数列的前10个数
	//for i := 0; i < 10; i++ {
	//	fmt.Println(fib())
	//}

	//var model ch5.Model
	//model.Name = "rain"
	//model.Age = 23
	//
	//fmt.Print(model.GetName())
	//fmt.Print(ch5.GetName(model))

	v := ch5.Vertex{3, 4}
	ch5.Scale(&v, 10)
	fmt.Println(ch5.Abs(v))
}
