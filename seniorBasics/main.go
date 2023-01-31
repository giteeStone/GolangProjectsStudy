package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

var (
	timeout int64
	size    int
	count   int
)

func main() {
	getCommandArgs()
	desIp := os.Args[len(os.Args)-1]
	conn, err := net.DialTimeout("ip:icmp", desIp, time.Duration(timeout)*time.Microsecond)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	data := make([]byte, size)

	fmt.Println(data, timeout, size, count)
}

func getCommandArgs() {
	flag.Int64Var(&timeout, "w", 1000, "请求超时时长，毫秒")
	flag.IntVar(&size, "l", 32, "请求发送缓冲区大小，字节")
	flag.IntVar(&count, "n", 4, "发送请求数")
	flag.Parse()
}
