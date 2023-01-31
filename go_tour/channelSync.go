package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done...")
	time.Sleep(time.Second)
	done <- true
}

func main() {

	done := make(chan bool, 1)
	go worker(done)

	msg := <-done //!!Block until we receive a notification from the worker on the channel.
	fmt.Print(msg)
}
