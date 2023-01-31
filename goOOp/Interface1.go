package main

import (
	"fmt"
)

// 定义一个接口
type USB interface {
	start()
	stop()
}

// 定义结构体
type Camera struct {
}

type Mouse struct {
}

// 实现接口方法
func (c *Camera) start() {
	fmt.Println("Camera started...")
}
func (c *Camera) stop() {
	fmt.Println("Camera stopped...")
}

func (m *Mouse) start() {
	fmt.Println("Mouse started...")
}
func (m *Mouse) stop() {
	fmt.Println("Mouse stopped...")
}

// 计算机的结构体
type Computer struct {
}

func (computer *Computer) Run(usb USB) {
	usb.start()
	usb.stop()
}

func main() {

	computer := &Computer{}
	m := &Mouse{}
	c := &Camera{}

	// 能够传入Run参数的是要求Mouse和Camera实现USB接口的所有方法
	// 如果Mouse和Camera结构体没有实现接口的所有方法会报错
	computer.Run(m)
	computer.Run(c)

}
