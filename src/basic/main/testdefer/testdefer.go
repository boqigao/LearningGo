package main

import "fmt"

func sum(a int, b int) int {
	// 当执行到defer时候，暂时不执行，将其压入独立的栈，defer栈
	// 当函数执行完毕之后，在从defer栈pop出来，先入后出
	defer fmt.Println("ok1 n1=", a)
	defer fmt.Println("ok2 n2=", b)
	res := a + b
	fmt.Println("ok3 res=", res)
	return res
}

func main() {
	res := sum(10, 20)
	fmt.Println("res=", res)
}
