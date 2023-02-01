// 在反射中关于类型还划分为两种：类型(Type)和种类(Kind)。因为在Go语言中我们可以使用type关键字构造很多自定义类型，而种类（Kind）就是指底层的类型，
// 但在反射中，当需要区分指针、结构体等大品种的类型时，就会用到种类（Kind）。举个例子，我们定义了两个指针类型和两个结构体类型，通过反射查看它们的类型和种类。
// 反射在运行期间，操作人员类型的对象。可以通过Typeof方法获得对象类型。通过ValueOf获取对象值。
package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func main() {
	stu1 := student{
		Name:  "小王子",
		Score: 90,
	}
	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind())
	//通过for循环遍历结构体的所有字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i) //求第i个字段
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	} //.NumField()返回字段数量

	//通过字段名获取指定结构体字段信息FieldByName,FieldByIndex
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}
}
