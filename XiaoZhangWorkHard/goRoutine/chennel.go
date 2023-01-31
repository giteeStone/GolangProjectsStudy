//channel

//var b chan <类型>
//通道必须用make函数初始化才能使用

// 通道的操作
// 1.发送：ch1<-1
// 2.接收：x:=<-ch1
// 3.关闭：close()
package main

import (
	"fmt"
	"sync"
)

var a []int
var b chan int //需要指定通道中元素的类型
var wg sync.WaitGroup

func noBufChannel() {
	b = make(chan int) //为通道分配内存，带缓冲区
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Println("后台goroutine从通道b中取到了", x)
	}()
	b <- 10
	fmt.Println("10发送到通道b中了...")
	wg.Wait()
	close(b)
}

func BufChannel() {
	b = make(chan int, 1) //为通道分配内存，不带缓冲区
	wg.Add(1)
	b <- 10 //因为缓冲区为1，所以只能放一个数据
	fmt.Println("10发送到通道b中了...")
	wg.Wait()
	close(b)
}

func main() {
	noBufChannel()
}
