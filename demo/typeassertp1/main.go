package main

import "fmt"

type Camera struct {
}

type Phone struct {
}

func (c Phone) Call() {
	fmt.Println("Phone call")
}

type Computer struct {
}

func (c Computer) Working(usb Usb) {
	usb.Start()
	// 类型断言
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
	usb.Stop()
}

// 声明 定义一个接口
type Usb interface {
	Start()
	// 如果指向 phone 还有 call 能力
	//Call()
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
	var usbArr [3]Usb
	usbArr[0] = Phone{}
	usbArr[1] = Phone{}
	usbArr[2] = Camera{}

	var computer Computer
	for _, v := range usbArr {
		computer.Working(v)
		fmt.Println("---------")
	}
}
