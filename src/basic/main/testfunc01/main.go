package main

import "fmt"

func getSum(a int, b int) int {
	return a + b
}

func myFunc(funVar func(int, int) int, a int, b int) int {
	return funVar(a, b)
}

func main() {
	//a := getSum
	//
	//fmt.Printf("a's type: %T\nfunc's type: %T\n", a, getSum)
	//
	//res := a(3, 5)
	//fmt.Println(res)

	res2 := myFunc(getSum, 8, 9)
	fmt.Println(res2)
}
