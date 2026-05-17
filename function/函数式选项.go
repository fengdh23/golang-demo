package main

import (
	"crypto/tls"
	. "fmt"
	"time"
)

/*
   耗子大佬 https://coolshell.cn/articles/21146.html
Go语言不支持重载函数,你得用不同的函数名来应对不同的参数的构造器
*/

type User struct {
	Id       string // 必须
	Name     string // 必须
	Age      int    // 非必须
	Gender   bool   // 非必须
	password string // 小写，非导出
}
type Server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxConns int
	TLS      *tls.Config
}

/*
非必输的选项都移到一个结构体里
*/
type Config struct {
	Protocol string
	Timeout  time.Duration
	Maxconns int
	TLS      *tls.Config
}

func main() {

	//for _, value := range os.Args {
	//	Println("Arg=", value)
	//}

	u := User{
		Id: "123", Name: "zhangsan", Age: 19, Gender: false, password: "123456",
	}
	Println(u)
}
