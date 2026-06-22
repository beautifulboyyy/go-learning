package main

import (
	"fmt"

	// 这是通过 go get 下载的第三方包，不属于 Go 标准库。
	"github.com/google/uuid"
)

func main() {
	// NewString 会生成一个新的 UUID 字符串，常用于唯一标识。
	id := uuid.NewString()
	fmt.Println("生成的 UUID：", id)
}
