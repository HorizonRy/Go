package main

import "fmt"

func main() {
	// 1. 声明一个空map
	// 创建一个map，键为string，值为int
	var myMap1 map[string]int
	if myMap1 == nil {
		fmt.Println("Map is empty")
	}

	// 分配空间
	myMap1 = make(map[string]int, 10)
	myMap1["one"] = 1
	myMap1["two"] = 2
	myMap1["three"] = 3

	fmt.Println(myMap1) // 打印出来为哈希顺序（乱序）
	for key, value := range myMap1 {
		fmt.Println(key, value)
	}

	fmt.Println("----------------------")

	// 2. 声明并初始化一个map
	myMap2 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println(myMap2)
	for key, value := range myMap2 {
		fmt.Println(key, value)
	}
}
