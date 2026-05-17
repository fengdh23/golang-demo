package main

import "fmt"

func main() {
	var i int = 100
	fmt.Println("i 的地址= ", &i)

	// 下面的 	var ptr *int = &i 的意思
	// 1. ptr 是一个指针变量
	// 2. ptr 变量的类型是 *int
	// 3. ptr 本身的值是 &i
	var ptr *int = &i

	fmt.Printf("ptr 存的值= %v\n", ptr)
	fmt.Printf("ptr 本身的地址=%v\n", &ptr)
	fmt.Printf("ptr 指向的值=%v\n", *ptr)

}
