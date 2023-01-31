package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	//定义模板，解析模板，渲染模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("err parse template")
		return
	}
	u1 := User{
		Name:   "小王子",
		gender: "男", //不可导出，模板不能读取
		Age:    18,
	}

	err = t.Execute(w, u1)

	if err != nil {
		fmt.Println("render template failed,err")
		return
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("err")

	}
}
