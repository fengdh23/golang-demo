package main

import "fmt"

// 返回类型是 函数
func scope() func() int {
	outer_var := 2
	foo := func() int { return outer_var }
	return foo
}

func outer() (func() int, int) {
	outer_var := 2
	inner := func() int {
		outer_var += 99
		return outer_var
	}
	inner()
	return inner, outer_var
}

func main() {

	// Outpus: 2
	//fmt.Println(scope()())

	inner, val := outer()
	fmt.Println(inner()) // => 200
	fmt.Println(val)     // => 101
}
