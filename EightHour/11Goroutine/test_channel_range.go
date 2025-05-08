package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}

		close(c)
	}()

	// channel c 关闭后退出循环
	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("main exit...")
}
