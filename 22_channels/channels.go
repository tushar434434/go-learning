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
/*
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
*/
//sync
//synchro

//buffer channel  sin    gle channel messaging

//
/*
func makeBuffer(size int) chan string {
	return make(chan string, size)
}

func main() {
	mails := makeBuffer(3)

	go func() {
		for i := 1; i <= 5; i++ {
			mails <- fmt.Sprintf("Mail %d", i)
		}
		close(mails)
	}()

	for mail := range mails {
		fmt.Println(mail)
	}
}*/

//multiple channels
func main() {
	email := make(chan string, 2)
	sms := make(chan string, 2)

	// Sending to email channel
	go func() {
		email <- "Email 1"
		email <- "Email 2"
		close(email)
	}()

	// Sending to SMS channel
	go func() {
		sms <- "SMS 1"
		sms <- "SMS 2"
		close(sms)
	}()

	// Receiving from email
	for msg := range email {
		fmt.Println("Email:", msg)
	}

	// Receiving from SMS
	for msg := range sms {
		fmt.Println("SMS:", msg)
	}
}
