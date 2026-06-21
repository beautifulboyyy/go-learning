package main

import (
	"fmt"
	"time"
)

// ==================== 3.3 Select 练习 ====================
//
// 核心概念：
//   select {
//   case msg := <-ch1:
//       // 处理 ch1 的数据
//   case msg := <-ch2:
//       // 处理 ch2 的数据
//   case <-time.After(time.Second):
//       // 超时处理
//   default:
//       // 没有数据时执行
//   }
//
// 对比 Python：
//   Python: select.select([ch1, ch2], [], [], timeout)
//   Go:     select { case <-ch1: ... case <-ch2: ... }
//
// ====================================================

// 练习1：基础 Select
// 同时等待两个 Channel
func testBasic() {
	ch1 := make(chan string, 1) // 缓冲 Channel
	ch2 := make(chan string, 1) // 缓冲 Channel

	// 启动两个 goroutine，分别向 ch1 和 ch2 发送数据
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "来自 ch1"
	}()

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch2 <- "来自 ch2"
	}()

	// TODO: 用 select 同时等待 ch1 和 ch2
	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case msg := <-ch2:
		fmt.Println(msg)
	}
}

// 练习2：超时处理
// 用 time.After 实现超时
func testTimeout() {
	ch := make(chan string, 1) // 缓冲 Channel

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch <- "数据"
	}()

	// TODO: 用 select 等待 ch，超时 100ms
	select {
	case msg := <-ch:
		fmt.Println("收到:", msg)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("超时")
	}

	// 等待 goroutine 完成
	time.Sleep(200 * time.Millisecond)
}

// 练习3：default 分支
// 没有数据时立即执行 default
func testDefault() {
	ch := make(chan int)

	// TODO: 用 select 检查 ch 是否有数据
	select {
	case msg := <-ch:
		fmt.Println("收到:", msg)
	default:
		fmt.Println("没有数据")
	}
}

// 练习4：循环 Select
// 持续监听多个 Channel
func testLoop() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(50 * time.Millisecond)
			ch1 <- fmt.Sprintf("ch1-%d", i)
		}
		close(ch1)
	}()

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(70 * time.Millisecond)
			ch2 <- fmt.Sprintf("ch2-%d", i)
		}
		close(ch2)
	}()

	// TODO: 用 for-select 循环接收数据，直到两个 Channel 都关闭
	for {
		// 先检查是否都关闭了
		if ch1 == nil && ch2 == nil {
			break
		}
		select {
		case msg, ok := <-ch1:
			if !ok {
				ch1 = nil
				continue
			}
			fmt.Println(msg)
		case msg, ok := <-ch2:
			if !ok {
				ch2 = nil
				continue
			}
			fmt.Println(msg)
		}
	}
}

func main() {
	fmt.Println("=== 练习1：基础 Select ===")
	testBasic()

	fmt.Println("\n=== 练习2：超时处理 ===")
	testTimeout()

	fmt.Println("\n=== 练习3：default 分支 ===")
	testDefault()

	fmt.Println("\n=== 练习4：循环 Select ===")
	testLoop()
}
