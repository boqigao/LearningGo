package main

import "fmt"

func main() {
	var slice = []int{1, 1}
	for i := 2; i < 10; i++ {
		slice = append(slice, slice[i-1]+slice[i-2])
	}
	fmt.Println(slice)
	
	fbnSlice := make([]int, 10)
	fbnSlice[0] = 1
	fbnSlice[1] = 1
	for i := 2; i < 10; i++ {
		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
	}
	fmt.Println(fbnSlice)
}
