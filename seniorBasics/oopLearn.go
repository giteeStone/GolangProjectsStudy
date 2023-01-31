package main

import (
	"fmt"
)

// 定义Student结构体和相关方法
type Student struct {
	name  string
	age   int
	score float32
}

func (stu *Student) ShowInfo() {
	fmt.Printf("学生名 = %v 年龄 = %v, 成绩 = %v\n", stu.name, stu.age, stu.score)
}

func (stu *Student) SetScore(score float32) {
	stu.score = score
}

// 小学生结构体
type Pupil struct {
	Student
}

// 结构体Pupil绑定的特有方法
func (p *Pupil) testing() {
	fmt.Println("小学生考试......")
}

// 大学生结构体
type Graduate struct {
	Student
}

// 结构体Graduate绑定的特有的方法
func (g *Graduate) testing() {
	fmt.Println("大学生测试......")
}

func main() {

	// 小学生
	puil := &Pupil{}
	puil.Student.name = "kiko"
	puil.Student.age = 8
	puil.Student.ShowInfo()
	puil.testing()
	puil.Student.SetScore(90)
	puil.Student.ShowInfo()

	fmt.Println("----------------------------------------------------")

	// 大学生
	graduate := &Graduate{}
	graduate.Student.name = "yoyo"
	graduate.Student.age = 18
	graduate.Student.ShowInfo()
	graduate.testing()
	graduate.Student.SetScore(130)
	graduate.Student.ShowInfo()
}
