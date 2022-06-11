package main

import "fmt"

func main() {
	num2 := new(int)
	//num2自己是一个指针，指向了一个地址，然后这个地址指向的实际值为0
	fmt.Printf("%T, %v, %v, %v \n", num2, num2, &num2, *num2)
	*num2 = 100
	fmt.Printf("%T, %v, %v, %v", num2, num2, &num2, *num2)
}
