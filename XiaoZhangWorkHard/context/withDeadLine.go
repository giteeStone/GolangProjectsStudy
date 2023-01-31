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
		case <-ctx.Done():
			return
		default:
		}
	}
}

func main() {
	d := time.Now().Add(1500 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	//func WithDeadline(parent Context,deadline time.Time)(Context,CancelFunc)
	//到了daedline这个时间点，parent就会过期自动

	//尽管ctx会过期，但在任何情况下调用它的cancel函数都是很好的实践
	//如果不这样做，可能会使上下文及其父类存活的时间超过必要的时间
	defer cancel()
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("周林")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

// func WithTimeout(parent Context,timeout Duration)(Context,CancelFunc)
// //取消此上下文将释放与其相关的资源，因此代码应该在此上下文中运行的

// 例：
// ctx,cancel:=context.WithTimeout(context.Background)(time.Millisecond*50)
// //时间到了就自动取消
