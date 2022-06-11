package main

import (
	"fmt"
	"strings"
)

// 其实就是构造了一个类，类里有个member叫suffix。。。搞得这么复杂
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	f := makeSuffix(".jpg")
	fmt.Println("wenjianming chulihou = ", f("winter"))
}
