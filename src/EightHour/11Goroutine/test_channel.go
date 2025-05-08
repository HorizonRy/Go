package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		defer fmt.Println("Goroutine 退出...")
		fmt.Println("Goroutine 正在运行...")

		c <- 666 // 使用无缓冲 channel 时，子 Goroutine 也会阻塞
	}()

	// 从 channel c 中读取数据
	num := <-c // channel c中没有数据会阻塞（线程间同步）
	fmt.Println("num = ", num)
	fmt.Println("main goroutine 退出...")
}
