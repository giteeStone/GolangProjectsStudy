package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	//b, err := ioutil.ReadFile("D:\\go_projects\\go\\go_prod\\webQiMi\\readFiles\\hello.txt")
	b, err := ioutil.ReadFile("./hello.txt")
	if err != nil {
		fmt.Print("file open error")
	}
	fmt.Println(string(b))
	fmt.Fprintln(w, string(b))
}

func main() {
	http.HandleFunc("/hello", hello)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Print("http serve failed,err")
		return
	}
}
