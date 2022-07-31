package main

import (
	"fmt"
	"net/http"
)

// 注意使用国内 url
var urls = []string{
	"http://www.baidu.com/",
	"http://golang.org/",
	"http://blog.golang.org/",
}

func main() {
	for _, url := range urls {
		resp, err := http.Head(url)
		if err != nil {
			fmt.Println("Error ", url, err)
		}
		fmt.Println(url, ": ", resp.Status)
	}
}
