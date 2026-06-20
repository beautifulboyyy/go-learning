package main

import "fmt"

// ==================== 2.1 指针练习 ====================
//
// 核心概念：
//   &变量   → 获取变量的地址
//   *指针   → 通过地址访问/修改变量的值
//   *类型   → 声明指针类型
//
// 对比 Python：
//   Python: def f(x): x = 100  （不会修改外部变量）
//   Go:     func f(p *int): *p = 100  （可以修改）
//
// ====================================================

// 练习1：创建指针
// 创建一个变量，然后创建一个指向它的指针，通过指针修改变量的值
func testPointer() {
	num := 42

	// TODO: 创建指针 p 指向 num
	// p := ???
	p := &num

	// TODO: 通过指针修改 num 的值为 100
	// *p = ???
	*p = 100

	fmt.Println("num =", num) // 应该输出 100
}

// 练习2：指针作为函数参数
// 实现 swap 函数，交换两个变量的值
func swap(a, b *int) {
	// TODO: 交换 a 和 b 指向的值
	tmp := *a
	*a = *b
	*b = tmp
}

func testSwap() {
	x, y := 10, 20
	swap(&x, &y)
	fmt.Println("x =", x, "y =", y) // 应该输出 x=20 y=10
}

// 练习3：new 函数
// 用 new 创建一个 *int，设置值并返回
func createPointer() *int {
	// TODO: 用 new 创建一个 *int
	// p := new(int)
	// 设置值为 999
	// 返回 p
	p := new(int)
	*p = 999
	return p
}

func testNew() {
	p := createPointer()
	fmt.Println("*p =", *p) // 应该输出 999
}

// 练习4：返回指针
// 实现 findMax 函数，返回较大值的指针
func findMax(a, b int) *int {
	// TODO: 返回较大值的指针
	// 提示：用 new(int) 创建，然后设置值
	p := new(int)
	if a > b {
		p = &a
	} else {
		p = &b
	}
	return p
}

func testFindMax() {
	p := findMax(10, 20)
	fmt.Println("max =", *p) // 应该输出 20
}

func main() {
	fmt.Println("=== 练习1：创建指针 ===")
	testPointer()

	fmt.Println("\n=== 练习2：指针作为参数 ===")
	testSwap()

	fmt.Println("\n=== 练习3：new 函数 ===")
	testNew()

	fmt.Println("\n=== 练习4：返回指针 ===")
	testFindMax()

	fmt.Println("\n✅ 全部完成！运行 go run ./basics/9-pointers/ 测试")
}
