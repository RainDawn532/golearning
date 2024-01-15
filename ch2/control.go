package ch2

import (
	"fmt"
	"runtime"
	"time"
)

// 流程控制相关
func ForFunc() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum = sum + i
	}
	fmt.Println(sum)
}

func SwitchFunc() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 2:
		fmt.Println("Tomorrow.")
	case today + 1:
		fmt.Println("In two days.")
	default:

		fmt.Println("Too far away.")
	}

}

/*
*Deferred function calls are executed in Last In First Out order after the surrounding function returns.
 */
func DeferFunc() (i int) {
	defer func() { i++ }()
	return 1
}
