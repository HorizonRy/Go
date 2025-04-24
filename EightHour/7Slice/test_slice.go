package main

import "fmt"

func main() {
	var myArray [10]int
	// 部分赋值，剩余元素为类型对应零值
	myArray2 := [10]int{1, 2, 3, 4}

	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}

	for index, value := range myArray2 {
		fmt.Println(index, ": ", value)
	}
}
