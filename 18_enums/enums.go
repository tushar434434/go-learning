package main

import "fmt"

//type OrderStatus int
type OrderStatus string

const (
	//received   OrderStatus = iota // iota is a predeclared identifier representing successive untyped integer constants (0, 1, 2, …) within a constant declaration.
	received   OrderStatus = "received"
	processing OrderStatus = "processing"
	shipped    OrderStatus = "shipped"
	delivered  OrderStatus = "delivered"
	canceled   OrderStatus = "canceled"
)

func changeStatus(status OrderStatus) {
	fmt.Println("Current status:", status)
}

// enumarated type
func main() {
	changeStatus(received)
}
