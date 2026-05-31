package main

import "fmt"

type Pupil struct {
	Name  string
	Age   int
	Score int
}

// 显示成绩
func (p *Pupil) ShowScore() {
	fmt.Println("Pupil:", p.Name, "Score:", p.Score)
}

// 打分
func (p *Pupil) SetScore(score int) {
	p.Score = score
}

func (p *Pupil) testing() {
	fmt.Println(p.Name + "小学生在考试")
}

type Graduate struct {
	Name  string
	Age   int
	Score int
}

func main() {
	// 测试 1
	var pupil = Pupil{"小王", 18, 10}
	pupil.testing()
	fmt.Println(pupil.Name)
}
