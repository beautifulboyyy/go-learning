package main

import "fmt"

// ==================== 2.5 defer/panic/recover 练习 ====================
//
// 核心概念：
//   defer  → 函数结束时执行（类似 Python 的 finally）
//   panic  → 程序崩溃（类似 Java 的 throw）
//   recover → 捕获 panic（类似 catch）
//
// 对比 Python：
//   Python: try: ... finally: ...  → Go: defer
//   Python: raise Exception(...)   → Go: panic(...)
//   Python: try: ... except: ...   → Go: defer + recover
//
// ====================================================

// 练习1：defer 基础
// 观察 defer 的执行顺序
func testDefer() {
	fmt.Println("开始")

	// TODO: 添加一个 defer，打印 "defer 1"
	// defer ???
	defer fmt.Println("defer 1")

	fmt.Println("中间")
	// TODO: 添加一个 defer，打印 "defer 2"
	// defer ???
	defer fmt.Println("defer 2")
	fmt.Println("结束")

	// 观察输出顺序：开始 → 中间 → 结束 → defer 2 → defer 1
}

// 练习2：defer 执行顺序
// defer 是后进先出（栈）
func testDeferOrder() {
	// TODO: 用 for 循环添加 3 个 defer，分别打印 1, 2, 3
	// for i := 1; i <= 3; i++ {
	//     defer fmt.Println(i)
	// }

	// 观察输出：3, 2, 1（后进先出）
	for i := 1; i <= 3; i++ {
		defer fmt.Println("defer: ", i)
	}
}

// 练习3：panic 和 recover
// 实现 safeDivide，用 recover 捕获 panic
func safeDivide(a, b int) int {
	// TODO: 用 defer + recover 捕获 panic
	// defer func() {
	//     if r := recover(); r != nil {
	//         fmt.Println("捕获到 panic:", r)
	//     }
	// }()

	// TODO: 如果 b == 0，触发 panic
	// if b == 0 { panic("除数不能为零") }
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("捕获到panic： ", r)
		}
	}()

	if b == 0 {
		panic("除数不能为0")
	}

	return a / b
}

func testPanicRecover() {
	result := safeDivide(10, 0)
	fmt.Println("结果:", result) // 即使 panic，程序也不会退出
}

// 练习4：defer 释放资源
// 模拟文件操作，用 defer 确保资源释放
func readFile(filename string) {
	// TODO: 打开文件（模拟）
	fmt.Println("打开文件:", filename)

	// TODO: 用 defer 关闭文件
	defer fmt.Println("关闭文件:", filename)

	// TODO: 读取文件内容（模拟）
	fmt.Println("读取文件内容")
}

func testResource() {
	readFile("test.txt")
	// 观察：打开 → 读取 → 关闭（defer 保证关闭一定执行）
}

func main() {
	fmt.Println("=== 练习1：defer 基础 ===")
	testDefer()

	fmt.Println("\n=== 练习2：defer 执行顺序 ===")
	testDeferOrder()

	fmt.Println("\n=== 练习3：panic 和 recover ===")
	testPanicRecover()

	fmt.Println("\n=== 练习4：defer 释放资源 ===")
	testResource()
}
