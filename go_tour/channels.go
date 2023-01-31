package main

import "fmt"

func main() {

	messages := make(chan string, 2)

	go func() {
		messages <- "ping"
		messages <- "hello"
	}()

	msg2 := <-messages
	msg1 := <-messages
	fmt.Println(msg2)
	fmt.Println(msg1)

}
