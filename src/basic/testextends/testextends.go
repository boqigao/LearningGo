package main

import "fmt"

type Goods struct {
	name  string
	price int
}

type Book struct {
	// 这种情况叫组合
	goods  Goods
	writer string
}

type Cookie struct {
	// 这种情况叫继承
	Goods
	size int
}

func main() {
	var book Book
	book.goods.price = 5
	book.goods.name = "#3"
	fmt.Println(book)
}
