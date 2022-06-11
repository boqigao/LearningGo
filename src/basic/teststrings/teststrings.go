package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "hello北"
	r := []rune(str)
	fmt.Println(len(str))
	fmt.Println(len(r))

	strInt := "12"
	n, err := strconv.ParseInt(strInt, 10, 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n)
	}
	var str3 string
	str3 = strconv.Itoa(12)
	fmt.Println(str3)

	var bytes = []byte("hello beijing")
	fmt.Printf("bytes=%v\n", bytes)

	str = string([]byte{97, 98, 99})
	fmt.Println(str)

	str = strconv.FormatInt(123, 16)
	fmt.Println(str)

	fmt.Println(strings.Contains("seafood", "afo"))

	fmt.Println(strings.Count("boboxika", "bo"))

	fmt.Println(strings.EqualFold("abc", "ABC"))

	fmt.Println(strings.Index("boboxika", "xika"))

	fmt.Println(strings.LastIndex("boboxikabobo", "bobo"))

	str4 := "go go hello"
	// str4本身不会变化
	str5 := strings.Replace(str4, "go", "北京", -1)
	fmt.Println(str5)

	strArr := strings.Split("hello,world,ok", ",")
	fmt.Println(strArr)

	str6 := strings.ToUpper(str4)
	fmt.Println(str6)

	var str7 string = "   abc f  "
	str8 := strings.TrimSpace(str7)
	fmt.Println(str8)

	var str9 string = " ! hello! "
	// 去掉cutset里面的所有字符
	str10 := strings.Trim(str9, "! ")
	fmt.Println(str10)
}
