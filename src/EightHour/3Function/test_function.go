package main

import "fmt"

func fool(a string, b int) int {
	fmt.Println("a = ", a, ", b = ", b)

	c := 100
	return c
}

// 匿名返回值
func fool2(a string, b int) (int, int) {
	fmt.Println("a = ", a, ", b = ", b)

	return 666, 777
}

// 有名返回值
func fool3(a string, b int) (r1, r2 int) {
	fmt.Println("a = ", a, ", b = ", b)

	r1, r2 = 888, 999
	return
}

func main() {
	c := fool("hello", 100)
	fmt.Println("c = ", c)

	ret1, ret2 := fool2("hello", 100)
	fmt.Println("ret1, ret2 = ", ret1, ret2)

	ret3, ret4 := fool3("hello", 100)
	fmt.Println("ret3, ret4 = ", ret3, ret4)
}
