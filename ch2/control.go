package ch2

import "fmt"

// 流程控制相关
func ForFunc() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum = sum + i
	}
	fmt.Println(sum)
}
