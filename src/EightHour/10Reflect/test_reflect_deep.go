package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Call() {
	fmt.Println("User Call()")
	fmt.Printf("User: %v\n", u)
}

func DoFiledAndMethod(input interface{}) {
	// 1. 获取input的type和value
	inputType := reflect.TypeOf(input)
	inputValue := reflect.ValueOf(input)
	fmt.Println(inputType)
	fmt.Println(inputValue)

	// 2. 获取input中的字段
	// 2.1 获取input的字段数量
	numField := inputType.NumField()
	fmt.Println("Field count:", numField)

	// 2.2 遍历每一个字段
	for i := 0; i < numField; i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("Field %d: %s, Value: %v\n", i, field.Name, value)
	}

	// 3. 获取input中的方法
	// 3.1 获取input的方法数量
	numMethod := inputType.NumMethod()
	fmt.Println("Method count:", numMethod)
	// 3.2 遍历每一个方法
	for i := 0; i < numMethod; i++ {
		method := inputType.Method(i)
		fmt.Printf("Method %d: %s\n", i, method.Name)
	}
}

func main() {
	user := User{Id: 1, Name: "Horizon", Age: 18}

	DoFiledAndMethod(user)

}
