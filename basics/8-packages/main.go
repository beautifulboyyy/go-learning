package main

import "learn-go/basics/8-packages/utils"

func main() {
	utils.SayHello("张三") // ✅ 可以调用（大写开头）
	// utils.secret()       // ❌ 不能调用（小写开头）
}
