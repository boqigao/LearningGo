package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// 对hero结构体切片进行排序

func main() {
	var intSlice = []int{0, -1, 10, 7, 90}
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄~%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heroes = append(heroes, hero)
	}

	for _, hero := range heroes {
		fmt.Println(hero)
	}
	fmt.Println("---------------------------")
	sort.Sort(heroes)

	for _, hero := range heroes {
		fmt.Println(hero)
	}
}

//声明hero结构体

type Hero struct {
	Name string
	Age  int
}

// 声明hero结构体的切片类型

type HeroSlice []Hero

// 实现Interface接口

// Len method
func (hs HeroSlice) Len() int {
	return len(hs)
}

// Less 按照hero的年龄从小到大排序
func (hs HeroSlice) Less(i int, j int) bool {
	// 记忆方法：如果i的年龄小于j的年龄为真，那么i排在j前面
	return hs[i].Age < hs[j].Age
}

func (hs HeroSlice) Swap(i int, j int) {
	//tmp := hs[i]
	//hs[i] = hs[j]
	//hs[j] = tmp

	// same  as
	hs[i], hs[j] = hs[j], hs[i]
}
