package main

import (
	"fmt"
)

type Monkey struct {
	name string
}

func (this *Monkey) climb() {
	fmt.Println(this.name, " 生来就会爬树...")
}

type FlyAble interface {
	canFly()
}
type SwimAble interface {
	canSwim()
}

type SmartMonkey struct {
	Monkey // 继承Monkey结构体
}

// 让SmartMonkey实现两个接口
func (this *SmartMonkey) canFly() {
	fmt.Println(this.name, " 是一个聪明的猴子，通过学习能够飞翔")
}
func (this *SmartMonkey) canSwim() {
	fmt.Println(this.name, " 是一个聪明的猴子，通过学习能够游泳")
}

func main() {

	monkey := SmartMonkey{
		Monkey{
			name: "yoyo",
		},
	}

	monkey.climb()
	monkey.canFly()
	monkey.canSwim()
}
