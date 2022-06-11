package main

import (
	"fmt"
)

func main() {
	var firstMap map[int]int
	firstMap = make(map[int]int, 19)
	firstMap[5] = 1
	firstMap[7] = 41
	fmt.Println(firstMap)

	// make理解成new就行了，分配地址
	secondMap := make(map[int]int)
	secondMap[1] = 3
	secondMap[2] = 3
	fmt.Println(secondMap)

	// map delete
	delete(secondMap, 1)
	fmt.Println(secondMap)
	delete(secondMap, 100)

	val, findRes := secondMap[2]

	fmt.Printf("%d\n", val)
	fmt.Println(findRes)
	fmt.Printf("%T\n", findRes)

	// map for range
	for k, v := range firstMap {
		println(k)
		println(v)
	}

	// map slice,其实就是list.add()
	var monsters []map[string]string
	monsters = make([]map[string]string, 2)

	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "bobo"
		monsters[0]["age"] = "100"
	}

	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "shika"
		monsters[1]["age"] = "200"
	}
	fmt.Println(monsters)
}
