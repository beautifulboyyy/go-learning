package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// Go 切片会被 Marshal 转换成 JSON 数组。
	users := []User{
		{Name: "小明", Age: 18},
		{Name: "小红", Age: 20},
	}

	data, err := json.Marshal(users)
	if err != nil {
		fmt.Println("转换 JSON 失败：", err)
		return
	}

	fmt.Println("切片转 JSON 数组：")
	fmt.Println(string(data))

	// JSON 数组会被 Unmarshal 解析成 Go 切片。
	input := []byte(`[{"name":"张三","age":21},{"name":"李四","age":22}]`)

	var decodedUsers []User
	// 传入切片的地址，Unmarshal 会创建元素并写入切片。
	err = json.Unmarshal(input, &decodedUsers)
	if err != nil {
		fmt.Println("解析 JSON 失败：", err)
		return
	}

	fmt.Println("JSON 数组转切片：")
	for _, user := range decodedUsers {
		fmt.Printf("姓名：%s，年龄：%d\n", user.Name, user.Age)
	}
}
