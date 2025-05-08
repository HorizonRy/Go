package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// 匿名函数形式创建goroutine
	go func(user string) {
		defer fmt.Printf("%s.defer", user)

		go func() {
			defer fmt.Println("B.defer")
			// return // 退出内层匿名函数
			runtime.Goexit() // 退出当前goroutine（直接退出外层函数）
			fmt.Println("B")
		}() // 调用匿名函数

		println("Hello", user)
	}("Horizon") // 匿名调用并传入参数

	for {
		time.Sleep(1 * time.Second)
	}
}
