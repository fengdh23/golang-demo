package main

import "fmt"

type Camera struct {
}

type Phone struct {
}

type Computer struct {
}

func (c Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}

// 声明 定义一个接口
type Usb interface {
	Start()
	Stop()
}

// 实现 USB 接口 的方法
func (c Phone) Start() {
	fmt.Println("phone start")
}
func (c Phone) Stop() {
	fmt.Println("phone Stop")
}

func (c Camera) Start() {
	fmt.Println("Camera Sart")
}
func (c Camera) Stop() {
	fmt.Println("Camera Stop")
}

func main() {
	computer := Computer{}
	//phone := Phone{}
	var camera = Camera{}

	//computer.Working(phone)
	computer.Working(camera)
}
