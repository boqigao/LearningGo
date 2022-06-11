package main

import "fmt"

func main() {

	//var intArr = [...]int{1, 22, 33, 66, 99}
	//slice := intArr[1:3] //下标为1到下标为3，不包含3，类似于substring
	//fmt.Println(slice)

	// slice内部维护了一个数组
	var slice2 = make([]int, 5, 10)
	fmt.Println(&slice2[0])

	var strs = []string{"boqi", "xika", "jiajia"}
	fmt.Println(strs)
	fmt.Println(len(strs))
	fmt.Println(cap(strs))

	strs = append(strs, "pipi")
	fmt.Println(strs)
	fmt.Println(len(strs))
	fmt.Println(cap(strs))

	var strs1 = []string{"popo"}
	strs = append(strs, strs1...)

	var slice3 = []int{1, 2, 3, 4}
	test(slice3)
	fmt.Println(slice3)

	// 修改字符串
	str := "boboxika"
	fmt.Println(&str)
	// 先转byte切片
	arr1 := []byte(str)
	arr1[2] = 'q'
	arr1[3] = 'i'
	// 再转回string，实际上以前的string地址变化了吗？并没有，为啥？
	// 因为通过切片修改了string内部的byte数组，所以并没有新建，卧槽
	str = string(arr1)
	fmt.Println(str)
	fmt.Println(&str)
}

func test(slice []int) {
	slice[0] = 100
}
