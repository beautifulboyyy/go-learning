package main

import (
	"encoding/json"
	"fmt"
)

// User 是 Go 中的用户结构体。
// JSON 标签用来指定转换后的 JSON 字段名。
type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	// 一、Marshal：把 Go 结构体转换成 JSON 字节。
	user := User{
		Name:  "小明",
		Age:   18,
		Email: "xiaoming@example.com",
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("转换 JSON 失败：", err)
		return
	}

	// jsonData 是 []byte，转换成 string 后方便阅读和打印。
	fmt.Println("结构体转 JSON：")
	fmt.Println(string(jsonData))

	// 二、Unmarshal：把 JSON 字节解析到 Go 结构体中。
	input := []byte(`{"name":"小红","age":20,"email":"xiaohong@example.com"}`)

	var decodedUser User
	// 必须传 &decodedUser，让 Unmarshal 能修改这个结构体。
	err = json.Unmarshal(input, &decodedUser)
	if err != nil {
		fmt.Println("解析 JSON 失败：", err)
		return
	}

	fmt.Println("JSON 转结构体：")
	fmt.Println("姓名：", decodedUser.Name)
	fmt.Println("年龄：", decodedUser.Age)
	fmt.Println("邮箱：", decodedUser.Email)
}
