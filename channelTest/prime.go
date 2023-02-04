package main

import (
	"fmt"
	"sync"
)

const (
	// 常量，表示协程的个数
	NUMROUNTINE = 10000
)

var (
	flag = NUMROUNTINE //标志协程个数，每一个协程都要经历退出循环操作

	locker sync.Mutex
)

// 协程函数，用于向numchan管道中写入数据
func writeNums(numchan chan int) {

	for i := 1; i <= cap(numchan); i++ {

		numchan <- i
	}
	// 写完则关闭管道
	close(numchan)
}

// 协程函数，供8个协程调用，读取numchan管道数据并且计算写入reschan管道
func getNums(numchan chan int, reschan chan map[int]int) {

	// 从numchan中读取数据
	for {
		val, ok := <-numchan
		if !ok { // 如果ok等于false，说明numchan被关闭了,且当前已经读完
			locker.Lock()
			flag-- // 这个是共享资源，必须要锁住,剩余需要完成计算求和的数字
			if flag == 0 {
				close(reschan)
			}
			locker.Unlock()
			break
		}
		// 计算结果
		m := make(map[int]int, 1)
		m[val] = getSum(val)
		reschan <- m
	}

}

func getSum(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}

func main() {

	// 创建两个管道
	var numchan chan int = make(chan int, 80000)
	var reschan chan map[int]int = make(chan map[int]int, 80000)

	go writeNums(numchan)

	for i := 0; i < NUMROUNTINE; i++ {
		go getNums(numchan, reschan)
	}

	for {
		val, ok := <-reschan
		if ok {
			fmt.Printf("%v\n", val)
		} else {
			break
		}
	}
}
