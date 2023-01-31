package main

//channel练习
//1.启动一个goroutine，生成100个数发送到ch1=a
//2.启动一个goroutine，从ch1中取值，计算其平方放到channel2中
//3.在main中，从ch2中取值

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var once sync.Once

func f1(ch1 chan int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

func f2(ch1, ch2 chan int) {
	defer wg.Done()
	for {
		x, ok := <-ch1 //ch1通道关闭后ok为false
		if !ok {
			break
		}
		ch2 <- x * x
	}
	once.Do(func() { close(ch2) }) //锁，确保某个操作只执行一次
}

func main() {
	a := make(chan int, 50) //一边放一边读，所以可以小于100,超过先阻塞
	b := make(chan int, 100)
	wg.Add(2)
	go f1(a)
	go f2(a, b)
	go f2(a, b)
	wg.Wait()
	for ret := range b {
		fmt.Println(ret)
	}
}

//单向通道

func f1(ch1 chan<- int) { //只能发送的通道类型
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}
func f2(ch1 chan<- int, ch2 <-chan int) { //只能发送和只能接受通道定义
	defer wg.Done()
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	close(ch1)
}
