package main

import (
	"fmt"
	"time"
)

// ==================== 3.1 Goroutine 练习 ====================
//
// 核心概念：
//   go 函数() → 启动一个 Goroutine 并发执行
//   Goroutine 是轻量级线程，由 Go 运行时调度
//
// 对比 Python：
//   Python: threading.Thread(target=func).start()
//   Go:     go func()
//
// ====================================================

// 练习1：基础 Goroutine
// 观察 Goroutine 的并发执行
func sayHello(name string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("%s: %d\n", name, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func testBasic() {
	// TODO: 用 go 启动 sayHello("goroutine")
	// go ???
	go sayHello("goroutine")
	// 主函数继续执行
	sayHello("main")

	// 观察：两个函数交替执行，而不是顺序执行
}

// 练习2：匿名 Goroutine
// 用匿名函数启动 Goroutine
func testAnonymous() {
	// TODO: 用 go 启动一个匿名函数，打印 "Hello from goroutine"
	// go func() {
	//     ???
	// }()
	go func() {
		fmt.Println("Hello from goroutine")
	}()

	time.Sleep(200 * time.Millisecond) // 等待 goroutine 执行
}

// 练习3：多个 Goroutine
// 同时启动多个 Goroutine
func testMultiple() {
	// TODO: 用 for 循环启动 3 个 Goroutine
	// for i := 1; i <= 3; i++ {
	//     go func(id int) {
	//         fmt.Printf("Goroutine %d\n", id)
	//     }(i)
	// }
	for i := 1; i <= 3; i++ {
		go func(id int) {
			fmt.Printf("Goroutine %d\n", id)
		}(i)
	}

	time.Sleep(200 * time.Millisecond) // 等待所有 goroutine 执行
}

// 练习4：观察并发执行
// 用 time.Sleep 观察 Goroutine 的执行顺序
func testConcurrent() {
	// TODO: 启动一个 Goroutine，打印 1-5
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Println("goroutine:", i)
			time.Sleep(50 * time.Millisecond)
		}
	}()

	// 主函数打印 1-5
	for i := 1; i <= 5; i++ {
		fmt.Println("main:", i)
		time.Sleep(50 * time.Millisecond)
	}
}

func main() {
	fmt.Println("=== 练习1：基础 Goroutine ===")
	testBasic()

	fmt.Println("\n=== 练习2：匿名 Goroutine ===")
	testAnonymous()

	fmt.Println("\n=== 练习3：多个 Goroutine ===")
	testMultiple()

	fmt.Println("\n=== 练习4：观察并发执行 ===")
	testConcurrent()

	fmt.Println("\n✅ 全部完成！")
}
