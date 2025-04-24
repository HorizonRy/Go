package main

import "fmt"

// 声明全局变量，不能使用 := 自动推断
var globalA int = 100
var globalB = 200

// globalC := 300

func main() {
	// 1. 单变量声明、初始化
	// 声明特定类型空变量
	var a int
	fmt.Println("a = ", a)

	// 声明特定类型且初始化
	var b int = 100
	fmt.Println("b = ", b)

	// 省去数据类型，自动推断
	var c = 100
	fmt.Println("c = ", c)

	// 【常用】省去 var，使用 := 初始化并自动推断
	// := 只适用于函数体内
	d := 100
	fmt.Println("d = ", d)
	fmt.Printf("type of d = %T\n", d)

	f := "abcd"
	fmt.Println("f = ", f)
	fmt.Printf("type of f = %T\n", f)

	g := 3.14
	fmt.Println("g = ", g)
	fmt.Printf("type of g = %T\n", g)

	// 尝试打印全局变量
	fmt.Println("globalA = ", globalA)
	fmt.Println("globalB = ", globalB)

	// 2. 多变量声明、初始化
	// 多变量为同种类型
	var xx, yy int = 100, 200
	fmt.Println("xx = ", xx, ", yy = ", yy)
	// 多变量为不同类型
	var kk, ll = 100, "ABCD"
	fmt.Println("kk = ", kk, ", ll = ", ll)
	// 多行形式
	var (
		vv int  = 100
		jj bool = true
	)
	fmt.Println("vv = ", vv, ", jj = ", jj)
}
