package main

import "fmt"

func addOne(a int) {
	a++
}

func main() {
	var a int = 3
	addOne(a)
	fmt.Println(a)
}
