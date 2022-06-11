package main

import "fmt"

type Point struct {
	x int
	y int
}

func TypeJudge(items ...interface{}) {
	for i, item := range items {
		// 一种固定用法
		switch item.(type) {
		case int:
			fmt.Println(i, "int")
		}
	}
}

func main() {
	var a interface{}
	var point Point = Point{1, 2}
	a = point

	b, ok := a.(int)

	if ok == true {
		fmt.Println("success")
		fmt.Println(b)
	} else {
		fmt.Println("fail")
	}
	// 不管是不是ok都能继续往下走
	fmt.Println("jixu")
}
