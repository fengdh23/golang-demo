package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point Point = Point{1, 2}
	a = point // 空接口 x 可以接受任何类型

	var b Point
	// b = a // 不可以
	b = a.(Point) // 类型断言
	fmt.Println(b)

	var x interface{}
	var b2 float32 = 1.1
	x = b2 // 空接口 x 可以接受任何类型

	//y, ok := x.(float64)
	if y, ok := x.(float32); ok {
		fmt.Println("类型断言ok")
		fmt.Printf("y的类型 %T, 值是 %v", y, y)
	} else {
		fmt.Println("类型断言失败")
	}

	fmt.Println("继续执行")
}
