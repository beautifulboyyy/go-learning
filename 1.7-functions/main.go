package main

import "fmt"

// 练习1：写一个函数，返回两个数的最大值
func max(a, b int) int {
	// 在这里写代码
	if a > b {
		return a
	} else {
		return b
	}
}

// 练习2：写一个函数，返回姓名和年龄（多返回值）
func getPerson() (string, int) {
	// 在这里写代码
	return "lsw", 26
}

// 练习3：写一个函数，计算任意个数的平均值
func average(nums ...float64) float64 {
	// 在这里写代码
	sum := 0.0
	for _, num := range nums {
		sum += num
	}
	return sum / float64(len(nums))
}

func main() {
	// 调用练习1
	fmt.Println(max(10, 20))

	// 调用练习2
	name, age := getPerson()
	fmt.Println(name, age)

	// 调用练习3
	fmt.Println(average(1, 2, 3, 4, 5))
}
