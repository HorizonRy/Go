package main

import "fmt"

// 使用 const 定义枚举类型
const (
	BEIJING  = 0
	SHANGHAI = 1
)

// 首个枚举元素使用 iota 实现自增编号
const (
	CHENGDU = iota
	KUNMING
)

func main() {
	const length int = 10
	fmt.Println(length)
	// 常量不允许修改
	// length = 100
}
