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
	name string
}

type Mouse struct {
	name string
}

// 实现接口方法
func (c *Camera) start() {
	fmt.Println(c.name, " Camera started...")
}
func (c *Camera) stop() {
	fmt.Println(c.name, " Camera stopped...")
}

func (m *Mouse) start() {
	fmt.Println(m.name, " Mouse started...")
}
func (m *Mouse) stop() {
	fmt.Println(m.name, " Mouse stopped...")
}

// 定义一个方法，且是Camera结构体自有的
func (c *Camera) snapshot() {
	fmt.Println(c.name, " Camera snapshot started...")
}

// 计算机的结构体
type Computer struct {
}

func (computer *Computer) Run(usb USB) {
	usb.start()
	// 类型断言
	// 如果能够转换为Camera，那么就执行它自有的结构体方法
	if obj, ok := usb.(*Camera); ok {
		obj.snapshot()
	}
	usb.stop()
}

func main() {

	computer := &Computer{}

	var usbArr [3]USB
	usbArr[0] = &Mouse{"极光鼠标"}
	usbArr[1] = &Camera{"太阳照相机"}
	usbArr[2] = &Mouse{"南极鼠标"}

	for i := 0; i < 3; i++ {
		computer.Run(usbArr[i])
	}

}
