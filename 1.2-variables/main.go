package main

import "fmt"

func main() {
	// 练习1：用短变量声明创建以下变量
	// 你的名字（string）
	// 你的年龄（int）
	// 你的身高（float64）
	var name string = "sw"
	var age int = 26
	var height float64 = 172.5

	fmt.Println(name, age, height)

	// 练习2：用 var 声明创建以下变量（不赋初始值）
	// 一个整数
	// 一个字符串
	// 一个布尔值
	// 然后打印出来看看零值是什么
	var num int
	var f float64
	var bol bool
	fmt.Println(num, f, bol)
	// 练习3：试试能不能这样做
	// a := 10
	// b := 3.14
	// c := a + b   ← 这样会报错吗？
	a := 10
	b := 3.14
	c := a + b
	fmt.Println(c)

}
