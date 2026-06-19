package main

import "fmt"

func main() {
	// 创建一个切片，包含 1, 2, 3
	// 添加 4, 5
	// 打印切片和长度
	var list []int = []int{1, 2, 3}
	list = append(list, 4, 5)
	fmt.Println(list, len(list))
}
