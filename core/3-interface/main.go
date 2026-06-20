package main

import "fmt"

// ==================== 2.3 Interface 练习 ====================
//
// 核心概念：
//   type 接口名 interface { 方法1() 返回类型; 方法2() 返回类型 }
//   实现：只要类型有接口要求的所有方法，就自动满足
//   使用：函数参数用接口类型，可以接受任何实现
//
// 对比 Java：
//   Java: class Dog implements Speaker  （显式声明）
//   Go:   type Dog struct{} + func (d Dog) Speak() {} （隐式满足）
//
// ====================================================

// 定义 Shape 接口
type Shape interface {
	Area() float64
}

// 定义 Rectangle 结构体
type Rectangle struct {
	Width  float64
	Height float64
}

// TODO: 为 Rectangle 实现 Area() 方法
// 提示：面积 = 宽 × 高
func (r Rectangle) Area() float64 {
	return r.Height * r.Width // 改成正确的计算
}

// 定义 Circle 结构体
type Circle struct {
	Radius float64
}

// TODO: 为 Circle 实现 Area() 方法
// 提示：面积 = π × 半径²，用 3.14159
func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius // 改成正确的计算
}

// 练习：实现 printArea 函数
// 接收 Shape 接口，打印面积
func printArea(s Shape) {
	// TODO: 打印 "面积: xxx"
	fmt.Println("面积：", s.Area())
}

func main() {
	fmt.Println("=== 练习1：实现接口方法 ===")

	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 3}

	fmt.Println("矩形面积:", rect.Area())   // 应该输出 50
	fmt.Println("圆形面积:", circle.Area()) // 应该输出 28.27431

	fmt.Println("\n=== 练习2：使用接口 ===")

	printArea(rect)   // 面积: 50
	printArea(circle) // 面积: 28.27431
}
