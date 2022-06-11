package main

type Person struct {
	Name string
}

// 这种情况下说明这个函数只能绑定类型A
func (a Person) test() {
	println(a.Name)
}

func main() {
	var person Person
	person.Name = "bobo"
	person.test()
}
