package main

import (
	"fmt"
	"time"
)

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time //nanosecond precision
}

// construtors hote nhi hai pr bna skete hai usning functions

func newOrder(id string, amount float32, status string, createdAt time.Time) order {
	return order{
		id:        id,
		amount:    amount,
		status:    status,
		createdAt: createdAt,
	}
}

func (o *order) changeStatus(status string) {
	o.status = status
}
func (o order) getAmount() float32 {
	return o.amount
}

func main() {
	order1 := newOrder("1", 50.00, "pending", time.Now())
	fmt.Println(order1)
	fmt.Println("order1 createdAt is:", order1.createdAt)
	fmt.Println("id is:", order1.id)
	order1.createdAt = time.Now()
	fmt.Println(order1)
	order2 := newOrder("2", 100.00, "shipped", time.Now())
	fmt.Println(order2)
	order1.changeStatus("delivered")
	fmt.Println(order1)
	fmt.Println("order1 amount is:", order1.getAmount())

	language := struct {
		name    string
		age     int
		student bool
	}{name: "Go", age: 10, student: true}
	fmt.Println(language)
}
