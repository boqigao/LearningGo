package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("%T\n%v\n", now, now)

	// 这几个数字一定要知道哈哈
	fmt.Println(now.Format("2006-01/02 15//04:05"))

	//i := 0
	//for {
	//	i++
	//	fmt.Println(i)
	//	// 不能去乘以0.1
	//	time.Sleep(100 * time.Microsecond)
	//	if i == 100 {
	//		break
	//	}
	//}

	//
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())

	cur := time.Now().Unix()
	test03()
	after := time.Now().Unix()

	fmt.Println(after - cur)

}

func test03() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)

	}
}