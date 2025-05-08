package main

import "fmt"

// 声明一种数据类型
type myInt int

// 定义一个结构体
type Book struct {
	title  string
	author string
}

func main() {
	// 使用自定义类型
	var a myInt = 10
	fmt.Printf("a = %d, type = %T\n", a, a)

	// 结构体使用
	var book1 Book
	book1.title = "Golang"
	book1.author = "Horizon"
	fmt.Printf("book1 = %v\n", book1)

	changeBook(book1)
	fmt.Printf("book1 = %v\n", book1)
}

func changeBook(book Book) {
	book.author = "Zry"
}
