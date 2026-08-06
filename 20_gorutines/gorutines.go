package main

import (
	"fmt"
	"time"
)

//	func task(id int) {
//		fmt.Println("doing task", id)
//	}
func main() {
	for i := 0; i <= 10; i++ {
		//go task(i)
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Second * 2)
} //for running proceses concurrently

/*Goroutines are lightweight threads managed by the Go runtime. You use them to run functions concurrently, handle asynchronous tasks, and build scalable programs without the heavy memory overhead of traditional operating system threads.*/
