package ch7

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

// tips  主协程执行完毕的话如果子协程还未执行的话也会退出执行
func printNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func printLetters() {
	for char := 'a'; char <= 'e'; char++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%c ", char)
	}
}

func run() {
	var wg sync.WaitGroup

	// 启动第一个协程
	wg.Add(1)
	go func() {
		defer wg.Done()
		printNumbers()
	}()

	// 启动第二个协程
	wg.Add(1)
	go func() {
		defer wg.Done()
		printLetters()
	}()

	// 等待两个协程完成
	wg.Wait()

	fmt.Println("\nMain goroutine exiting.")
}

// chan  for goroutine
func forSum(ins []int, s chan int) {
	sum := 0
	for _, value := range ins {
		sum = sum + value
	}
	s <- sum
}

func RunWithChan() {
	//信道
	c := make(chan int)

	s := []int{1, 2, 32, 21}

	go forSum(s[:len(s)/2], c)
	go forSum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从 c 中接收

	fmt.Println(x, y, x+y)
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// C chan 一边需要不断读取  一边迟迟不放入数据  导致  fatal error: all goroutines are asleep - deadlock!  死锁  互相等待  必须close()信道
func C() {

	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}

	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	// c2 <- "hello"

	select {
	case msg1 := <-c1:
		fmt.Println("c1 received: ", msg1)
	case msg2 := <-c2:
		fmt.Println("c2 received: ", msg2)
	default:
		fmt.Println("default :s")
	}
}

//等价二叉查找树查找

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
	len   int
}

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *Tree, ch chan int) {
	if nil != t.Left {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if nil != t.Right {
		Walk(t.Right, ch)
	}
}
func Same(t1, t2 *Tree) bool {
	treeValue1 := make(chan int)
	treeValue2 := make(chan int)
	go Walk(t1, treeValue1)
	go Walk(t2, treeValue2)
	//
	for i := 0; i < t1.len; i++ {
		x := <-treeValue1
		y := <-treeValue2
		if x != y {
			return false
		}
	}
	return true
}

// sync.Mutex进行同步

// SafeCounter 的并发使用是安全的。
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc 增加给定 key 的计数器的值。
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	c.v[key]++
	c.mux.Unlock()
}

// Value 返回给定 key 的计数器的当前值。
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	defer c.mux.Unlock()
	return c.v[key]
}

///todo 泛型

func ToJSON[T any](v T) string {
	b, _ := json.Marshal(v)
	return string(b)
}
