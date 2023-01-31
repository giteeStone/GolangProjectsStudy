package main

import (
	"flag"
	"fmt"
	_ "fmt"
)

var (
	timeout int64
	size    int
	count   int
)

func main() {

	getCommandArgs()
	fmt.Println(timeout, size, count)
}

func getCommandArgs() {
	flag.Int64Var(&timeout, "w", 1000, "请求超时时长，毫秒")
	flag.IntVar(&size, "l", 32, "请求发送缓冲区大小，字节")
	flag.IntVar(&count, "n", 4, "发送请求数")
	flag.Parse()
}
