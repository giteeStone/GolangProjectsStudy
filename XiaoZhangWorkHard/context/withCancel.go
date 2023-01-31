// 为啥需要context  go 1.7
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func f(ctx context.Context) {
	defer wg.Done()
	for {
		fmt.Println("**")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done(): //Done在select中使用，withCancel对应的当cancel（）调用时返回的channel关闭
			return
		default:
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background()) //造一个取消的函数,Background()造一个根节点 ,子节点是ctx(WithCancel类型)
	wg.Add(1)
	go f(ctx) //传递上下文，main 线程可以取消子线程，通过ctx通讯
	time.Sleep(time.Second * 5)
	cancel() //调用cancel造的函数，通知goroutine退出
	wg.Wait()

}
