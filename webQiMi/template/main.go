package main

import (
	"fmt"
	"html/template"
	"net/http"
)

//遇事不决写注释

func sayHello(w http.ResponseWriter, r *http.Request) {
	//解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("parse template err ")
		return
	}

	//渲染模板
	err = t.Execute(w, "小王子")
	if err != nil {
		fmt.Println("render template failed,err")
		return
	}
}
func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("Http server start failed")
		return
	}
}
