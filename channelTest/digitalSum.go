/*使用 goroutine 和 channel 实现一个计算int64随机数各位数和的程序，例如生成随机数61345，计算其每个位数上的数字之和为19。
开启一个 goroutine 循环生成int64类型的随机数，发送到jobChan
开启24个 goroutine 从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
主 goroutine 从resultChan取出结果并打印到终端输出*/

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func writeData(w chan uint64, wg *sync.WaitGroup) {
	for i := 0; i < 800; i++ {
		w <- rand.Uint64()
	}
	close(w)
	wg.Done()
}

func procData(w chan uint64, r chan map[uint64]uint64, wg *sync.WaitGroup) {
	for {
		i64, ok := <-w
		if !ok {
			wg.Done()
			break
		} else {
			newmap := make(map[uint64]uint64, 1)
			newmap[i64] = getSumUint64(i64)
			r <- newmap
		}
	}
}

func getSumUint64(sums uint64) uint64 {
	var total uint64 = 0
	fmt.Println(sums)
	for {
		if sums == 0 {

			break
		} else {
			total += sums % 10
			sums = sums / 10
		}
	}
	fmt.Println(total)
	return total
}

func main() {

	wchan := make(chan uint64, 800)
	sumChan := make(chan map[uint64]uint64, 800)
	wg.Add(1)
	go writeData(wchan, &wg)

	intProcs := runtime.NumCPU()
	runtime.GOMAXPROCS(intProcs)
	fmt.Println("cores ", intProcs)
	for i := 0; i < intProcs; i++ {
		wg.Add(1)
		go procData(wchan, sumChan, &wg)
	}
	time.Sleep(1 * time.Second)
	for {
		k, ok := <-sumChan
		if ok {
			fmt.Printf("%v\n", k)
		} else {
			break
		}
	}
	wg.Wait()
}
