package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age,omitempty"`
	Email string `json:"email,omitempty"`
}

func main() {
	// Age 和 Email 没有赋值，所以它们分别是 0 和空字符串。
	user := User{
		Name: "小明",
	}

	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("转换 JSON 失败：", err)
		return
	}

	// 因为 Age 和 Email 带有 omitempty，空值字段不会出现在 JSON 中。
	fmt.Println(string(data))

	// 字段有值时，即使带有 omitempty，也会正常出现在 JSON 中。
	completeUser := User{
		Name:  "小红",
		Age:   20,
		Email: "xiaohong@example.com",
	}

	completeData, err := json.Marshal(completeUser)
	if err != nil {
		fmt.Println("转换 JSON 失败：", err)
		return
	}

	fmt.Println(string(completeData))
}
