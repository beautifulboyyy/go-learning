package main

import (
	"errors"
	"fmt"
)

// ==================== 2.4 错误处理练习 ====================
//
// 核心概念：
//   函数返回 (值, error)
//   出错时：return 0, errors.New("错误信息")
//   成功时：return 值, nil
//   调用时：if err != nil { 处理错误 }
//
// 对比 Python：
//   Python: try: result = divide(10, 0) except: print("出错")
//   Go:     result, err := divide(10, 0); if err != nil { ... }
//
// ====================================================

// 练习1：返回错误
// 实现 divide 函数，除数为 0 时返回错误
func divide(a, b float64) (float64, error) {
	// TODO: 如果 b == 0，返回错误
	// 提示：errors.New("除数不能为零")
	if b == 0 {
		return 0, errors.New("除数不能为零")
	}
	return a / b, nil
}

func testDivide() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("错误:", err) // 应该输出：错误: 除数不能为零
		return
	}
	fmt.Println("结果:", result)
}

// 练习2：带格式的错误
// 实现 validateAge 函数，年龄必须在 0-150 之间
func validateAge(age int) error {
	// TODO: 如果 age < 0 或 age > 150，返回错误
	// 提示：fmt.Errorf("年龄无效: %d", age)
	if age < 0 || age > 150 {
		return fmt.Errorf("年龄无效: %d", age)
	}
	return nil
}

func testValidate() {
	err := validateAge(200)
	if err != nil {
		fmt.Println("验证失败:", err) // 应该输出：验证失败: 年龄无效: 200
	}
}

// 练习3：多返回值处理
// 实现 parsePerson 函数，返回姓名和年龄
func parsePerson(name string, ageStr string) (string, int, error) {
	// TODO: 如果 name 为空，返回错误
	// TODO: 简单处理：ageStr 必须是数字（这里简化，只要不为空就行）
	// 提示：返回 "", 0, err 或 name, age, nil
	if name == "" {
		return "", 0, errors.New("name为空")
	}
	if ageStr == "" {
		return "", 0, errors.New("age为空")
	}
	return name, 0, nil
}

func testParse() {
	name, age, err := parsePerson("", "25")
	if err != nil {
		fmt.Println("解析失败:", err)
		return
	}
	fmt.Printf("姓名: %s, 年龄: %d\n", name, age)
}

// 练习4：错误链
// 实现 readFile 函数，模拟读取文件出错
func readFile(filename string) (string, error) {
	// TODO: 如果 filename 为空，返回错误
	// 否则返回 "文件内容"
	if filename == "" {
		return "", errors.New("filename为空")
	}
	return "", nil
}

func testReadFile() {
	content, err := readFile("")
	if err != nil {
		fmt.Println("读取失败:", err)
		return
	}
	fmt.Println("内容:", content)
}

func main() {
	fmt.Println("=== 练习1：返回错误 ===")
	testDivide()

	fmt.Println("\n=== 练习2：验证错误 ===")
	testValidate()

	fmt.Println("\n=== 练习3：多返回值 ===")
	testParse()

	fmt.Println("\n=== 练习4：错误链 ===")
	testReadFile()
}
