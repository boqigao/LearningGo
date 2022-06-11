package main

import "fmt"

type MethodUtils struct {
}

func (receiver MethodUtils) Draw() {
	str := ""
	for i := 0; i < 8; i++ {
		str += "*"
	}
	for i := 0; i < 10; i++ {
		fmt.Println(str)
	}
}

func main() {
	var util MethodUtils

	util.Draw()
}
