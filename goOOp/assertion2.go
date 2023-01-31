package main

import "fmt"

type Student struct {
	name string
	age  int
}

func TypeJudge(items ...interface{}) {

	for index, val := range items {

		switch val.(type) { // 固定写法type是关键字

		case bool:
			fmt.Printf("第%v个参数是bool类型，值是%v\n", index, val)

		case float32:
			fmt.Printf("第%v个参数是float32类型，值是%v\n", index, val)

		case float64:
			fmt.Printf("第%v个参数是 float64类型，值是%v\n", index, val)

		case int, int32, int64:
			fmt.Printf("第%v个参数是整数类型，值是%v\n", index, val)

		case string:
			fmt.Printf("第%v个参数是string类型，值是%v\n", index, val)

		case Student:
			fmt.Printf("第%v个参数是student类型，值是%v\n", index, val)

		case *Student:
			fmt.Printf("第%v个参数是*student类型，值是%v\n", index, val)

		default:
			fmt.Printf("第%v个参数是类型不确定，值是%v\n", index, val)
		}
	}
}

func main() {

	TypeJudge(100, true, 34.9, "str", Student{name: "yoyo", age: 18}, &Student{name: "kiko", age: 19})

}
