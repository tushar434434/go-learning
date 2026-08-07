package main

import "fmt"

func main() {
	/*	messagechan := make(chan string)
		messagechan <- "ping" //channels are blocking
		msg := <-messagechan
		fmt.Println(msg)*/
	ch := make(chan string) // create channel
	// sender goroutine
	go func() {
		ch <- "Hello" // send
	}()
	msg := <-ch // receive
	fmt.Println(msg)

}
