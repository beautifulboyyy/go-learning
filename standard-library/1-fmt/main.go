package main

import "fmt"

type Product struct {
	Name  string
	Price float64
	Stock int
}

func main() {
	product := Product{
		Name:  "Go 入门书",
		Price: 59.9,
		Stock: 8,
	}

	fmt.Printf("商品：%s，价格：%.2f，库存：%d\n", product.Name, product.Price, product.Stock)

	fmt.Printf("product类型：%T\n", product)
	fmt.Printf("product：%+v\n", product)
	msg := fmt.Sprintf("%s × 2 = %.2f 元", product.Name, product.Price*2)
	fmt.Println(msg)
}
