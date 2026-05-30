package main

import "fmt"

type Box struct {
	len    float64
	width  float64
	height float64
}

// 方法 获取体积
func (b *Box) getVolume() float64 {
	return b.len * b.width * b.height
}

// 景区门票案例 不同年纪不同价格

type Visitor struct {
	Name string
	Age  int
}

func (visitor *Visitor) showPrice() {
	// 年龄大小范围验证
	if visitor.Age < 0 || visitor.Age > 100 {
		fmt.Println("visitor.Name, 0")
		return
	}
	if visitor.Age > 18 {
		fmt.Printf("visitor.Name %v, 年龄 %v 门票 18\n", visitor.Name, visitor.Age)
	} else {
		fmt.Printf("visitor.Name %v, 年龄 %v 门票 免费\n", visitor.Name, visitor.Age)
	}
}

func main() {
	var box Box
	box.len = 5
	box.width = 3
	box.height = 2

	volume := box.getVolume()
	fmt.Printf("体积是: %v \n", volume)

	var v Visitor
	for {
		fmt.Println("请输入 visitor.Name:")
		if v.Name == "n" {
			fmt.Println("结束")
			break
		}
		fmt.Scanln(&v.Name)
		fmt.Println("请输入 visitor.Age:")
		fmt.Scanln(&v.Age)

		v.showPrice()
	}
}
