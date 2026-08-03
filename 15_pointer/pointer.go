package main

import "fmt"

func change(num *int) {
	*num = 1
	fmt.Println("value of num is:", *num)
}
func main() {
	a := 10
	fmt.Println("value of a is:", a)
	change(&a)
	fmt.Println("value of a is:", a)

}
