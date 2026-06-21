package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	sourcePath := "standard-library/4-files/notes.txt"
	targetPath := "standard-library/4-files/notes-copy.txt"

	source, err1 := os.Open(sourcePath)
	if err1 != nil {
		fmt.Println("文件打开失败:", err1)
		return
	}
	defer source.Close()

	target, err2 := os.Create(targetPath)
	if err2 != nil {
		fmt.Println("文件创建失败:", err2)
		return
	}
	defer target.Close()

	written, err3 := io.Copy(target, source)
	if err3 != nil {
		fmt.Println("文件复制出错:", err3)
		return
	}
	fmt.Println("完成复制：", written)
}
