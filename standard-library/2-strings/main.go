package main

import (
	"fmt"
	"strings"
)

func main() {
	raw := "  Go,Python,Java  "

	// TODO 1：使用 TrimSpace 去掉 raw 两端的空白。
	cleaned := strings.TrimSpace(raw)

	// TODO 2：使用 Split 按逗号拆分，得到 []string。
	items := strings.Split(cleaned, ",")

	// TODO 3：使用 Contains 检查清理后的字符串是否包含 "Python"。
	havePython := strings.Contains(cleaned, "Python")
	// TODO 4：使用 ToLower 把清理后的字符串转换成小写。
	lower := strings.ToLower(cleaned)
	// TODO 5：使用 Join 将切片以 " | " 连接起来。
	join := strings.Join(items, "|")
	// 按下面格式输出结果：
	// 清理后：Go,Python,Java
	// 切片：[Go Python Java]
	// 包含 Python：true
	// 小写：go,python,java
	// 重新连接：Go | Python | Java
	fmt.Println(cleaned)
	fmt.Println(items)
	fmt.Println(havePython)
	fmt.Println(lower)
	fmt.Println(join)
	fmt.Println("请完成 TODO")
}
