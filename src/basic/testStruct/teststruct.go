package main

import (
	"encoding/json"
	"fmt"
	"unsafe"
)

type Cat struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Color string `json:"color"`
}

func main() {
	var cat1 Cat
	cat1.Name = "xiaobai"
	cat1.Color = "baise"
	cat1.Age = 3

	// struct是一个值类型，这两个地址相同
	fmt.Printf("%p\n", &cat1)
	fmt.Println(&cat1.Name)
	//Strings in Go are represented by reflect.
	//StringHeader containing a pointer to actual string data and a length of string:
	fmt.Println(unsafe.Sizeof(cat1.Name))
	fmt.Println(unsafe.Sizeof(&cat1.Name))

	jsonMonster, err := json.Marshal(cat1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("jsonStr", string(jsonMonster))
	}
}
