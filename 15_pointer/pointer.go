package main

import "fmt"

func change(num *int) {//passing the address of a to the function
	*num = 1//dereferencing the pointer to change the value of a
	fmt.Println("value of num is:", *num)
}
func main() {
	a := 10
	fmt.Println("value of a is:", a)
	change(&a)
	fmt.Println("value of a is:", a)

}
