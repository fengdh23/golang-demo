# golang-util


Study Golang.
1. 基础语法
2. 一些小demo.

类型推导

var p = xx --》  p := xx

包


在 import 包时，路径从$GOPATH的src下开始，不用带src,编译器会自动从src下开始引入



变量 到 指针

类型分类：值类型和引用类型
1. 值类型：
   2. 基本数据类型，比如 int、float、string、bool，还有数组和结构体 struct
   3. 变量直接存储值，内存（值）通常在栈中分配
2. 引用类型：
   3. 指针、Map、切片 slice 管道chan、接口 interface。
   4. 变量存储的是地址，这个地址对应的空间才真正存储数据的值。内存（值）通常在堆中分配。
   

函数（不同与方法）

  func 韩树明(入参) 返回{

}

函数 function 是独立的：


方法 method 是绑定到某个类型上的函数，多了一个 接收者 receiver：
```go
type User struct {
Name string
}

func (u User) SayHello() {
fmt.Println("Hello,", u.Name)
}
```

