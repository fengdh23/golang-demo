package main

// 结构体 小写 不可导出，又期望其他包也创建这个结构体实例
type student struct {
	Name string
	age  int // 如果小写，其它包不可访问.
}

// 前提是这个类可以修改,通过方法把所有的改成公有的
func (s *student) GetScore() int {
	return s.age
}

// 前提是这里可以修改,通过方法把所有的改成公有的
// 工厂模式来创建结构体实例 转换下

func NewStudent(name string, age int) *student {
	return &student{name, age}
}
