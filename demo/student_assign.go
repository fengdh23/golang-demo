package main

import "fmt"

type Stu struct {
	Name string
	Age  int
}

func main() {
	// 方式 1 结构体创建变量时，直接赋值 依赖字段定义的顺序
	var stu1 = Stu{"flynn", 38}
	stu2 := Stu{"hai", 23}

	// 方式 2 字段名和字段值写在一起 不依赖字段定义的顺序
	var stu3 = Stu{
		Age:  20,
		Name: "jack",
	}

	fmt.Println(stu1, stu2, stu3)

	// 返回结构体的指针类型

	var stu11 = &Stu{"flynn", 29}
	var stu22 *Stu = &Stu{"flynn", 29}
	fmt.Println(stu11)  // 地址 &{flynn 29}
	fmt.Println(*stu11) // {flynn 29}
	fmt.Println(*stu22) // {flynn 29}

}
