package main

import (
	"fmt"
	"sync"
)

// 三个函数分别输出“dog""fish""cat",交替轮流执行100次
func dog(dogch, catch chan struct{}, wg *sync.WaitGroup) {
	for {
		if dogCounter <= 0 {
			wg.Done()
			return
		}
		<-catch
		fmt.Println("dog", dogCounter)
		dogCounter--
		dogch <- struct{}{}
	}
}

func fish(fishch, dogch chan struct{}, wg *sync.WaitGroup) {
	for {
		if fishCounter <= 0 {
			wg.Done()
			return
		}
		<-dogch
		fmt.Println("fish", fishCounter)
		fishCounter--
		fishch <- struct{}{}
	}
}

func cat(catch, fishch chan struct{}, wg *sync.WaitGroup) {
	for {
		if catCounter <= 0 {
			wg.Done()
			return
		}
		<-fishch
		fmt.Println("cat", catCounter)
		catCounter--
		catch <- struct{}{}
	}
}

var dogCounter int8 = 100
var fishCounter int8 = 100
var catCounter int8 = 100

func main() {
	//次数
	var wg sync.WaitGroup
	dogChan := make(chan struct{}, 1)
	fishChan := make(chan struct{}, 1)
	catChan := make(chan struct{}, 1)
	wg.Add(3)
	catChan <- struct{}{}
	go dog(dogChan, catChan, &wg)
	go fish(fishChan, dogChan, &wg)
	go cat(catChan, fishChan, &wg)
	wg.Wait()
}
