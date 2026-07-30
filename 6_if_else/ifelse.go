package main

import "fmt"

func main() {
	age := 16
	if age >= 18 {
		fmt.Println("person is adult")
	} else {
		fmt.Println("Person is not an adult")
	}

	student := 100
	if student >= 120 {
		fmt.Println("over")
	} else if student >= 100 {
		fmt.Println("okay strength")
	} else {
		fmt.Println("less")
	}
	//go does not have ternary, you will  have to use normal if else

	if num := 130; num >= 18 {
		fmt.Println("person is adult")
	} else if num >= 12 {
		fmt.Println("person is teen")
	}
}
