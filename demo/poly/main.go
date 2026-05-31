package main

import "fmt"

type Camera struct {
	Name string
}

type Phone struct {
	Name string
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
	// 定义一个 Usb 接口数组，可以存放 Phone 和 Camera
	// 实现多态数组 -- 因为数组数据类型是一样的，通过接口可以存放 Phone 和 Camera
	// Phone 和 camera 可能能力不一样,需要类型断言
	var usbArr [3]Usb
	fmt.Printf("usbArr:%T\n", usbArr)
	// 创建 Phone 实例，并赋值给 Usb 接口数组
	usbArr[0] = Phone{"vivo"}
	usbArr[1] = Phone{" iphone"}
	usbArr[2] = Camera{"手机摄像头"}

	fmt.Println(usbArr)
}
