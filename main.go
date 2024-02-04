package main

import "golearning/ch5"

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

	//v := ch5.Vertex{3, 4}
	//p := &v
	//fmt.Println(ch5.Abs(v))
	//
	//q := p
	//
	//ch5.Scale(p)
	//ch5.Scale1(v)
	//ch5.Scale(q)
	//fmt.Println("指针修改值")
	//fmt.Println(p)
	//fmt.Println(q)
	//fmt.Println(v)

	//person := ch6.Person{"lnh", 23, "qwer1234"}
	//fmt.Println(person.Fly("ccd"))
	//plane := ch6.AirPlane{"飞机"}
	//
	//fmt.Println(plane.Fly("滑行"))
	//
	//fmt.Printf("(%v, %T)\n", plane, plane)
	//
	//var i interface{}
	//fmt.Printf("(%v, %T)\n", i, i)
	//i = 42
	//fmt.Printf("(%v, %T)\n", i, i)
	//
	//ch6.Do(person)
	//
	//fmt.Println(person.String())
	//
	//addr := ch6.IPAddr{1, 2, 3, 4}
	//fmt.Println(addr)

	//i, err := strconv.Atoi(";")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("Converted integer:", i)

	//sqrt, err := ch6.Sqrt(-2)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("Converted float:", sqrt)

	//ch6.ReaderFromFile()
	//ch6.WriteToFile()

	//go func() {
	//	fmt.Println("1231")
	//}()
	//
	//time.Sleep(1 * time.Second)
	//
	//ch7.RunWithChan()
	//
	//ch7.C()

	//ch8.Func4Reflect(ch5.Model{Age: 23, Name: "lnh"})

	//ch9.RunWithDelay()
	//ch8.GetHttp()
	//ch4.MapFunc()

	var model ch5.Model
	//model := ch5.Model{1, "lnh"}

	ch5.MethodSet(model) // 显示T方法集
	println("---------------")
	ch5.MethodSet(&model) // 显示*T方法集
}
