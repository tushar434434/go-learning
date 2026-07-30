package main

import (
	"fmt"
	"time"
)

func main() {
	//simple
	i := 3
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("other")
	}
	//multiple condition switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("its a weekday")
	default:
		fmt.Println("its a workkday")
	}

	whoamI := func(i interface{}) {
		switch k := i.(type) {
		case int:
			fmt.Println("its a integer")
		case string:
			fmt.Println("its a string")
		case bool:
			fmt.Println("boolean")

		default:
			fmt.Println("other", k)

		}
	}

	whoamI(50)

}
