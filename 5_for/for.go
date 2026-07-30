package main

import "fmt"

//for -> only construct in go for looping
func main() {
	//while loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	//classic for loop
	for j := 0; j <= 3; j++ {
		if j == 2 {
			continue
		}
		fmt.Println(j)
	}
	for l := range 11 {
		fmt.Println(l)
	}
}
