package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3)

	// 查看 channel 缓冲区的长度和容量
	fmt.Println("len:", len(c), "cap:", cap(c))

	go func() {
		defer fmt.Println("子 goroutine 结束...")

		for i := 0; i < 4; i++ {
			c <- i // channel c 缓冲区满时才发生阻塞，否则不阻塞
			fmt.Printf("子 goroutine 发送数据：%d, len = %d, cap = %d\n", i, len(c), cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++ {
		num := <-c // channel c 缓冲区为空时才发生阻塞，否则不阻塞
		fmt.Println("接收数据：", num)
	}

	time.Sleep(1 * time.Second)

	fmt.Println("main goroutine 结束...")
}
