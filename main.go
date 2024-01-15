package main

import "golearning/ch4"

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
	ch4.G()
}
