package main

import "fmt"

func add(nums ...int) int { // we can pass interface for any time of arguments
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}
func main() {
	fmt.Println(add(1, 2, 3, 4, 5))
}
