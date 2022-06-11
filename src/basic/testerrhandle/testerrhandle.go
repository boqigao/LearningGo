package main

import (
	"errors"
	"fmt"
)

func main() {
	test()
	fmt.Println("main ximian daima")

	test02()
}

func test() {
	defer func() {
		err := recover() // recover是内置函数，可以捕获异常
		if err != nil {
			fmt.Println("err=", err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println(res)
}

// 函数去读取一个配置文件的信息
// 如果文件名传入不正确，我们就返回一个自定义的错误
func readConf(name string) (err error) {
	if name == "conf.ini" {
		return nil
	} else {
		return errors.New("fail to read file")
	}
}

func test02() {
	err := readConf("confff.ini")
	if err != nil {
		panic(err)
	}

	fmt.Println("test02 continue")
}
