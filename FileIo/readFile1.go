//文件操作：

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// 文件读取
func readbybyte(filename string) {
	fileObj, err := os.Open(filename)
	//fileObj是*os.File类型，err是error类型
	if err != nil {
		fmt.Printf("open file failed, err:%v", err)
		return
	}
	//记得关闭文件
	defer fileObj.Close() //当err不为nil时，fileObj为nil，不能调用Close方法
	//读文件
	//var tmp =make([]byte,128) 指定读的长度
	var tmp [128]byte
	for {
		n, err := fileObj.Read(tmp[:])
		if err == io.EOF || n <= 0 {
			return
		}
		if err != nil {
			fmt.Printf("read from file failed, err:%v", err)
			return
		}
		fmt.Printf("读了%d个字节\n", n)
		fmt.Println(string(tmp[:n]))
	}
}

// bufio按行读取示例
func readbybufio(filename string) {
	fileObj, err := os.Open(filename)
	if err != nil {
		fmt.Printf("open file failed, err:%v", err)
		return
	}
	defer fileObj.Close()
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n') //注意是字符
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read from file failed, err:%v", err)
			return
		}
		fmt.Println(line)
	}
}

// 按ioutil一次读取全部文件
func readbyioutil(filename string) {
	ret, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("read file failed,err:%v\n", err)
		return
	}
	fmt.Println(string(ret))
}

// 文件写入
func writefile(filename string) {
	fileObj, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND, 0644) //有则追加，无则创建
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return
	}
	fileObj.Write([]byte("zhoulin mengbi le!"))
	fileObj.WriteString("周林解释不了")
	fileObj.Close()
}

func writefile2(filename string) {
	fileObj, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return
	}
	defer fileObj.Close()
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("hello沙河\n") //写到缓存中
	wr.Flush()                  //把缓存中的内容写入文件
}

func writefile3(filename string) {
	str := "hello 沙河"
	err := ioutil.WriteFile(filename, []byte(str), 0666)
	if err != nil {
		fmt.Printf("read file failed,err:%v\n", err)
		return
	}
}

func fileinsert(filename string) {
	fileObj, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("文件打开失败")
		return
	}
	defer fileObj.Close()
	fileObj.Seek(4, 0) //光标从0号字节开始移动一个字节,注意换行符算两个字节
	var ret [1]byte
	n, err := fileObj.Read(ret[:])
	if err != nil {
		fmt.Println("读取失败")
		return
	}
	fmt.Println(string(ret[:n]))
	fileObj.Seek(1, 0)
	var s []byte
	s = []byte{'c'}
	fileObj.Write(s) //会把原文件相应位置的字符换成c
	//如果在不改变原文件的情况下插入，需要将原文件在插入位置截断，然后将前半部分，插入字符和后半部分到新文件
}

// bufio获取用户的输入
func useBufio() {
	var s string
	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	fmt.Printf("你输入的内容是：%s\n", s) //可以在输入中接受空格
}

// 输出重定向
func main() {
	fileObj, err := os.OpenFile("./xx.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v", err)
		return
	}
	log.SetOutput(fileObj) //将log输出指定为文件fileObj
	for {
		log.Println("这是一条测试的日志")
		time.Sleep(time.Second * 3)
	}
}
