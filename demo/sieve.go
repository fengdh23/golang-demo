// Package main 实现了一个基于管道的素数筛算法
// 该程序使用并发的方式生成和过滤素数,展示了 Go 语言 channel 的使用
package main

import "fmt"

// Generate 函数向通道中持续发送从2开始的递增整数序列
// 参数:
//   - ch: 只写通道,用于发送生成的整数
func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

// Filter 函数从输入通道接收整数,过滤掉能被指定素数整除的数
// 将不能被整除的数发送到输出通道,实现素数筛选的核心逻辑
// 参数:
//   - in: 只读通道,接收待过滤的整数
//   - out: 只写通道,发送过滤后的整数
//   - prime: 用于判断的素数,如果能被整除则丢弃
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

// main 函数是程序的入口点
// 通过串联多个 Filter 协程实现素数筛算法:
// 1. 启动 Generate 协程生成初始数字流
// 2. 循环10次,每次从通道取出第一个数(必定是素数)
// 3. 为每个找到的素数创建新的 Filter 协程,过滤掉其倍数
// 4. 通过重新赋值 ch = ch1 将过滤器串联起来
func main() {
	ch := make(chan int)
	go generate(ch)
	for i := 0; i < 10; i++ {
		prime := <-ch
		fmt.Println(prime)
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
}
