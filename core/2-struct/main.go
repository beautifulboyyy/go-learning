package main

import "fmt"

// 定义 Book 结构体
type Book struct {
	Title  string
	Author string
}

// ==================== 2.2 Struct 练习 ====================
//
// 核心概念：
//   type 名字 struct { 字段1 类型; 字段2 类型 }
//   创建：Person{Name: "张三", Age: 25}
//   访问：p.Name
//   修改：p.Name = "李四"
//
// 对比 Python：
//   Python: class Person: def __init__(self, name, age)
//   Go:     type Person struct { Name string; Age int }
//
// ====================================================

// 练习1：定义并创建 Struct
// 定义一个 Book 结构体，包含 Title 和 Author 字段
func testBook() {
	// TODO: 创建一个 Book，Title="Go语言", Author="张三"
	// book := Book{???}
	book := Book{"Go语言", "张三"}

	fmt.Println(book) // {Go语言 张三}
}

// 练习2：修改字段
// 创建后修改字段值
func testModify() {
	book := Book{Title: "Python入门", Author: "李四"}

	// TODO: 把 Author 改成 "王五"
	// ??? = ???
	book.Author = "王五"
	fmt.Println(book) // {Python入门 王五}
}

// 练习3：Struct 作为函数参数
// 实现 printBook 函数，打印书的信息
func printBook(b Book) {
	// TODO: 打印 "书名: xxx, 作者: xxx"
	fmt.Println("书名：", b.Title, "作者：", b.Author)
}

func testPrint() {
	book := Book{Title: "Go实战", Author: "赵六"}
	printBook(book)
}

// 练习4：返回 Struct
// 实现 createBook 函数，返回一个新的 Book
func createBook(title, author string) Book {
	// TODO: 创建并返回一个 Book
	return Book{title, author}
}

func testCreate() {
	book := createBook("算法导论", "CLRS")
	fmt.Println(book) // {算法导论 CLRS}
}

func main() {
	fmt.Println("=== 练习1：创建 Struct ===")
	testBook()

	fmt.Println("\n=== 练习2：修改字段 ===")
	testModify()

	fmt.Println("\n=== 练习3：Struct 作为参数 ===")
	testPrint()

	fmt.Println("\n=== 练习4：返回 Struct ===")
	testCreate()
}
