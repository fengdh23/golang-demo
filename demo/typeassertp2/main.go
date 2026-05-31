package main

import "fmt"

func TypeJudge(items ...interface{}) {
	for index, x := range items {
		switch x.(type) { //这里的 type 是一个关键词，固定写法
		case int, int32, int64:
			fmt.Printf("第%v 个参数是 int 类型，值是%v \n", index, x)
		case string:
			fmt.Printf("第%v 个参数是 string 类型，值是%v \n", index, x)
		case bool:
			fmt.Printf("第%v 个参数是 bool 类型，值是%v \n", index, x)
		case float32:
			fmt.Printf("第%v 个参数是 float32 类型，值是%v \n", index, x)
		case float64:
			fmt.Printf("第%v 个参数是 float64 类型，值是%v \n", index, x)
		default:
			fmt.Println("other")
		}
	}
}

func main() {
	var n1 float32 = 1.1
	TypeJudge(1, "hello", true, n1, 1.2)
}
