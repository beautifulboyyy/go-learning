package main

import "fmt"

func main() {
	// 练习1：判断一个数是正数、负数还是零
	num := -5
	// 在这里写 if 判断
	if num > 0 {
		fmt.Println("正数")
	} else if num < 0 {
		fmt.Println("负数")
	} else {
		fmt.Println("0")
	}

	// 练习2：用 for 打印 1-10
	// 在这里写 for 循环
	for i := 1; i < 11; i++ {
		fmt.Print(i)
	}
	fmt.Println()
	// 练习3：用 switch 判断今天是周几
	day := "Wednesday"
	// 在这里写 switch
	switch day {
	case "Monday":
		fmt.Println("周一")
	case "Tuesday":
		fmt.Println("周二")
	case "Wednesday":
		fmt.Println("周三")
	default:
		fmt.Println("other")
	}
}
