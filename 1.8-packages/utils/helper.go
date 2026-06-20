package utils

import "fmt"

// 大写开头 = 公开
func SayHello(name string) {
	fmt.Println("Hello,", name)
}

// 小写开头 = 私有
func secret() {
	fmt.Println("这是私有函数")
}
