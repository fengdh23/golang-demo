package main

import (
	"fmt"
	"time"
)

func main() {
	result := 0
	start := time.Now()
	//                             YYYY MM DD HH mm ss
	dateStr := start.Format("2006_01_02_15_04_05")
	fmt.Println(dateStr) // 2022_09_03_17_34_40
	nowUTC := time.Now().UTC()
	fmt.Println(nowUTC.Local().Local()) //2022-09-03 17:38:36.9055364 +0800 CST

	//
	for i := 0; i <= 41; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf(" took this amount of time : %s", delta)
}
func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}
