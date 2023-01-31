package main

import "fmt"

type A struct {
	Name string
	age  int
}

func (a *A) SayOk() {
	fmt.Println("A SayOk ", a.Name)
}

func (a *A) SayHello() {
	fmt.Println("A SayHello ", a.Name)
}

type B struct {
	A
}

func main() {

	var b B
	b.A.Name = "kiko"
	b.A.age = 10
	b.A.SayOk()
	b.A.SayHello()

	// 上面的写法可以简化
	b.Name = "yoyo"
	b.age = 20
	b.SayOk()
	b.SayHello()
}
