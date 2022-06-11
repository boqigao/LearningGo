package main

import "fmt"

type Usb interface {
	// Start 生命了两个没有实现的方法
	Start()
	Stop()
}

// Phone Golang中不需要显式的实现，只要含有接口中的所有方法，那么这个struct就实现了这个接口
type Phone struct {
}

type Camera struct {
}

func (p Phone) Start() {
	fmt.Println("phone start")
}

func (p Phone) Stop() {
	fmt.Println("phone stop")
}

func (p Camera) Start() {
	fmt.Println("camera start")
}

func (p Camera) Stop() {
	fmt.Println("camera stop")
}

type Computer struct {
}

// Working 只要是实现了usb接口的所有方法，就是实现了usb接口
func (c Computer) Working(usb Usb) {
	// 通过usb接口来调用
	usb.Start()
	usb.Stop()
}

func main() {
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	computer.Working(phone)
	computer.Working(camera)
}
