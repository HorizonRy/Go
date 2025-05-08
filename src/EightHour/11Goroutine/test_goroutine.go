package main

import (
	"fmt"
	"time"
)

func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("newGoroutine: i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// 创建一个新的 goroutine，执行newTask函数
	go newTask()

	i := 0

	for {
		i++
		fmt.Printf("main Goroutine: i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}
