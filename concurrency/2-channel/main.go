package main

import "fmt"

// ==================== 3.2 Channel 练习 ====================
//
// 核心概念：
//   ch := make(chan int)      → 创建无缓冲 Channel
//   ch := make(chan int, 10)  → 创建有缓冲 Channel
//   ch <- 42                 → 发送数据
//   value := <-ch            → 接收数据
//   close(ch)                → 关闭 Channel
//
// 对比 Python：
//   Python: queue.Queue()    → 线程安全的队列
//   Go:     make(chan int)   → goroutine 间通信
//
// ====================================================

// 练习1：基础 Channel
// 用 Channel 在两个 goroutine 之间传递数据
func testBasic() {
	// TODO: 创建一个无缓冲 Channel
	ch := make(chan int)

	// TODO: 启动一个 goroutine，发送数据 42
	go func() {
		ch <- 42
	}()

	// TODO: 主 goroutine 接收数据
	value := <-ch
	fmt.Println("接收到:", value)
}

// 练习2：有缓冲 Channel
// 观察有缓冲 Channel 的行为
func testBuffered() {
	// TODO: 创建一个容量为 3 的 Channel
	ch := make(chan int, 3)

	// TODO: 发送 3 个数据（不会阻塞）
	ch <- 1
	ch <- 2
	ch <- 3

	// TODO: 接收 3 个数据
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// 练习3：Channel 作为函数参数
// 实现 producer-consumer 模式
func producer(ch chan int) {
	// TODO: 发送 1-5 到 Channel
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	// TODO: 关闭 Channel
	close(ch)
}

func consumer(ch chan int) {
	// TODO: 从 Channel 接收数据，直到关闭
	for value := range ch {
		fmt.Println("消费:", value)
	}
}

func testProducerConsumer() {
	// TODO: 创建 Channel，启动 producer 和 consumer
	ch := make(chan int)
	go producer(ch)
	consumer(ch)
}

// 练习4：多个 goroutine 通信
// 用 Channel 实现多个 goroutine 协作
func worker(id int, ch chan string) {
	// TODO: 发送工作完成消息
	ch <- fmt.Sprintf("Worker %d 完成", id)
}

func testMultiple() {
	// TODO: 创建 Channel，启动多个 worker
	ch := make(chan string)
	for i := 1; i <= 3; i++ {
		go worker(i, ch)
	}
	// TODO: 接收所有消息
	for i := 0; i < 3; i++ {
		fmt.Println(<-ch)
	}
}

func main() {
	fmt.Println("=== 练习1：基础 Channel ===")
	testBasic()

	fmt.Println("\n=== 练习2：有缓冲 Channel ===")
	testBuffered()

	fmt.Println("\n=== 练习3：Producer-Consumer ===")
	testProducerConsumer()

	fmt.Println("\n=== 练习4：多个 goroutine 通信 ===")
	testMultiple()
}
