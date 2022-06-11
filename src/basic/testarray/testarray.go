package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var hens [6]float32
	hens[0] = 3.0
	hens[1] = 3.0
	hens[2] = 3.0
	hens[3] = 3.0
	hens[4] = 3.0
	hens[5] = 3.0

	fmt.Printf("%p\n", &hens)
	fmt.Println(&hens[0])
	fmt.Println(&hens[1])
	fmt.Println(unsafe.Sizeof(hens[1]))

	var numArr01 = [3]int{2, 3, 4}
	fmt.Println(numArr01)

	var numArr02 = [3]int{2, 3, 1}
	fmt.Println(numArr02)

	var numArr03 = [...]int{2, 3, 4, 1}
	fmt.Println(numArr03)

	var numArr04 = [...]int{1: 900, 0: 22, 2: 321}
	fmt.Println(numArr04)

	for index, value := range numArr04 {
		fmt.Println(index)
		fmt.Println(value)
	}

	numArr05 := [...]int{}
	fmt.Println(numArr05)

	var arr [3]int

	fmt.Println(&arr[0])

	test01(arr)
	fmt.Println(arr)

	test02(&arr)
	fmt.Println(arr)
}

// 在java里属于引用传递！！在这里是值传递
func test01(arr [3]int) {
	arr[0] += 10
	fmt.Println(&arr[0])
}

// 传个地址来
func test02(arr *[3]int) {
	arr[0] += 10
	fmt.Println(&arr[0])
}
