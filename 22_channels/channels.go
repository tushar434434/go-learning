package main

import "fmt"

// func main() {
// 	/*	messagechan := make(chan string)
// 		messagechan <- "ping" //channels are blocking
// 		msg := <-messagechan
// 		fmt.Println(msg)*/
// 		ch := make(chan string) // create channel
// 	// sender goroutine
// 	go func() {
// 		ch <- "Hello" // send
// 	}()
// 	msg := <-ch // receive
// 	fmt.Println(msg)

// }

//synchronization using channels======
//sync
func worker(done chan bool) {
	fmt.Println("Working...")
	done <- true // signal completion
}
func main() {
	done := make(chan bool)
	go worker(done)
	<-done // wait for signal (sync point)
	fmt.Println("Done!")
}

//sync
//synchro
