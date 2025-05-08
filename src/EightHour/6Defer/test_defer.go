package main

import "fmt"

func deferFunc() int {
	// 顺序入栈，逆序执行
	fmt.Println("deferFunc2")
	fmt.Println("deferFunc1")
	return 0
}

func returnFunc() int {
	// return 的执行优先于 defer
	fmt.Println("returnFunc")
	return 0
}

func deferReturnFunc() int {
	defer deferFunc()
	return returnFunc()
}

func main() {
	deferReturnFunc()
}
