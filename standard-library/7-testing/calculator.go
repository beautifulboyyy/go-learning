package calculator

import "errors"

// Add 返回两个整数的和。
func Add(a, b int) int {
	return a + b
}

// Divide 执行整数除法，除数为零时返回错误。
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("除数不能为零")
	}

	return a / b, nil
}
