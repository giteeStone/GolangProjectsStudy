package main

import (
	"fmt"
	"sync"
)

var Gorouts, flag int = 10000, 10000
var locker sync.Mutex

func writeData(inchan chan int) {
	for i := 1; i < cap(inchan); i++ {
		inchan <- i
	}
	close(inchan)
}

func readDataPrime(inchan chan int, sumChan chan map[int]int) {
	for {
		num, ok := <-inchan
		if !ok {
			locker.Lock()
			flag--
			if flag == 0 {
				close(sumChan)

			}
			locker.Unlock()
			break
		}
		map1 := make(map[int]int, 1)
		map1[num] = getSum(num)
		sumChan <- map1
	}
}

func getSum(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i
	}
	return sum
}
func main() {
	//写入80000个素数
	inchan := make(chan int, 80000)
	sumChan := make(chan map[int]int, 80000)
	go writeData(inchan)
	//读取inchan并判断素数1万线程读8万数字
	for i := 1; i <= Gorouts; i++ {
		go readDataPrime(inchan, sumChan)
	}

	//读出sumChan
	for {
		map1, ok := <-sumChan
		if ok {
			fmt.Printf("%v\n", map1)
		} else {
			break
		}

	}
}
