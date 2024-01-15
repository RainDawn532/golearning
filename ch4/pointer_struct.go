package ch4

import "fmt"

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
	//
	ints := make([]int, 0, 5)
	fmt.Println(len(ints), cap(ints))
}
