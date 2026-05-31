package main

// 悟空像猴子一样会爬树，也有自己的能力，会飞
import "fmt"

// Monkey 结构体
type Monkey struct {
	Name string
}

// 爬树
func (m *Monkey) ClimbTree() {
	fmt.Println("生来会爬树")
}

// LitteMonkey 悟空 结构体
type LittleMonkey struct {
	Monkey // 继承
}

type BirdAble interface {
	Flying()
}

type FishAble interface {
	Swimming()
}

// LittleMonkey 实现 BirdAble 接口
func (this *LittleMonkey) Flying() {
	fmt.Println(this.Name, "会飞")
}

// LittleMonkey 实现 FishAble 接口
func (this *LittleMonkey) Swimming() {
	fmt.Println(this.Name, "会游泳")
}

func main() {
	monkey := LittleMonkey{
		Monkey{
			Name: "悟空",
		}, // ，不能少
	}
	monkey.ClimbTree()
	monkey.Flying()
	monkey.Swimming()

}
