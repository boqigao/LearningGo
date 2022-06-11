package main

import (
	"firstProject/src/basic/testfactory/model"
	"fmt"
)

func main() {
	var stu = model.NewStudent("tom~", 99)
	fmt.Println(*stu)
}
