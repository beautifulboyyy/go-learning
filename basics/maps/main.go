package main

import "fmt"

func main() {
	m := map[string]string{
		"name": "张三",
		"city": "北京",
	}

	// 练习1：获取 name，接两个值
	// 用 realName 和 exists 作为变量名

	// 练习2：获取一个不存在的 key "age"，看看 exists2 是什么

	// 练习3：遍历打印所有键值对
	// 用 for k, v := range m
	realName, exists := m["name"]
	if exists {
		fmt.Println(realName)
	}
	age, ok := m["age"]
	if ok {
		fmt.Println(age)
	}

	for k, v := range m {
		println(k, v)
	}

}
