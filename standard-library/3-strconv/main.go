package main

import (
	"fmt"
	"strconv"
)

func main() {
	ageText := "18"
	priceText := "59.9"
	activeText := "true"

	age, err := strconv.Atoi(ageText)
	if err != nil {
		fmt.Println("输入不是数字:", err)
		return
	}
	fmt.Println("年龄：", age)

	price, err2 := strconv.ParseFloat(priceText, 64)
	if err2 != nil {
		fmt.Println("浮点解析错误:", err2)
		return
	}
	fmt.Printf("价格：%.2f\n", price)

	active, err3 := strconv.ParseBool(activeText)
	if err3 != nil {
		fmt.Println("bool解析出错:", err3)
		return
	}
	fmt.Println("是否启用：", active)

	age += 1
	ageText = strconv.Itoa(age)
	fmt.Printf("明年年龄：%s,类型：%T\n", ageText, ageText)
}
