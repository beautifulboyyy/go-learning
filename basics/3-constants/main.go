package main

import "fmt"

const APP_NAME = "我的应用"

const (
	Small  = iota // 0
	Medium        // 1
	Large         // 2
)

func main() {
	fmt.Println(APP_NAME)
	fmt.Println(Small, Medium, Large)
}
