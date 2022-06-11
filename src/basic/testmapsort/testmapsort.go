package main

import (
	"fmt"
	"sort"
)

func main() {
	mapSort := make(map[int]int)

	for i := 0; i < 20; i++ {
		mapSort[i] = i
	}

	// 直接打印的话的确是无序的
	for k, v := range mapSort {
		fmt.Println(k, v)
	}

	// 遍历key
	var keys []int
	for k, _ := range mapSort {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	fmt.Println(keys)

	for _, k := range keys {
		fmt.Println(mapSort[k])
	}

}
