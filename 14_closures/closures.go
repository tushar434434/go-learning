package main

import "fmt"

func counter() func() int {
	var count int = 0
	return func() int {
		count++
		return count
	}
}

func main() {
	c := counter()
	fmt.Println(c()) // Output: 1
	fmt.Println(c()) // Output: 2
	fmt.Println(c()) // Output: 3
}
