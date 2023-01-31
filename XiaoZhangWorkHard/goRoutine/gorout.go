package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func f() {
	defer wg.Done() //执行完wg减1
	for i := 0; i < 10; i++ {
		r1 := rand.Int()    //int64的随机数
		r2 := rand.Intn(10) //[0,10)
		fmt.Println(r1, r2)
	}
}

var wg sync.WaitGroup //类似于os中的信号量

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1) //每次执行wg加1
		go f()
	}
	wg.Wait() //等wg的计数器减为0时结束
}
