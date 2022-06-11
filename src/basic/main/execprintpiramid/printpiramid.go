package main

import "fmt"

func main() {
	var n int = 20
	// 传入了地址
	test(&n)

	fmt.Println(n)
}

// 错误：收到了地址，然后取这个地址的值？ （把* 和 int分开？）
// 正确：n1 本身是一个指针类型，然后呢，在这个函数内部用 *n1来取到的n1的值
// 直接改变了n1指向的空间的值，~
func test(n1 *int) {
	*n1 = *n1 + 10
}
