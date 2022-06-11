package main

import "fmt"

type Monkey struct {
	Name string
}

func (m Monkey) climbing() {
	fmt.Println(m.Name, "can climbing")
}

type BirdAble interface {
	Flying()
}

type LittleMonkey struct {
	Monkey
}

func (lm LittleMonkey) Flying() {
	fmt.Println(lm.Name, "try to fly")
}

func main() {
	
}
