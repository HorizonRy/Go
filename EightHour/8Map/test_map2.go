package main

import "fmt"

func main() {
	cityMap := map[string]string{
		"China": "Beijing",
		"USA":   "NewYork",
		"Japan": "Tokyo",
	}

	// 添加元素
	cityMap["Germany"] = "Berlin"

	// 遍历
	printMap(cityMap)

	// 删除元素
	delete(cityMap, "Japan")

	// 判断元素是否存在
	if value, ok := cityMap["Japan"]; ok {
		fmt.Println("Japan:", value)
	} else {
		fmt.Println("Japan: not found")
	}

	// 修改元素
	changeValue(cityMap, "USA", "Washington")
	changeValue(cityMap, "UK", "London")

	for key, value := range cityMap {
		fmt.Println(key, value)
	}
}

func printMap(cityMap map[string]string) {
	for key, value := range cityMap {
		fmt.Println(key, value)
	}
	fmt.Println("------------------------")
}

func changeValue(cityMap map[string]string, key string, value string) {
	cityMap[key] = value
}
