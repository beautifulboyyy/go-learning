package main

import (
	"fmt"
	"os"
)

func main() {
	path := "standard-library/4-files/notes.txt"
	content := "Go 文件操作\n第二行内容。"

	// 文件写入的是字节，所以把 string 转成 []byte。
	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		fmt.Println("写入失败:", err)
		return
	}

	// ReadFile 一次读取整个文件，返回 []byte。
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("读取失败:", err)
		return
	}

	// 确定字节是文本后，将 []byte 转回 string。
	fmt.Println(string(data))
}
