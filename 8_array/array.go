package main

import "fmt"

func main() {
	var nums [4]int
	nums[0] = 10
	nums[1] = 20
	nums[2] = 30
	nums[3] = 40
	fmt.Println(len(nums))
	fmt.Println(nums[0])
	var n [3]bool
	fmt.Println(n)
	// int mein hogi 0 bool mein hogi false string mein hogi khali string
	//to declare in singlr line
	num := [3]int{1, 2, 3}
	fmt.Println(num)

	//2d array
	twoD := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(twoD)
	//fixed size agr go
	//memory optimization
	//constant time exces
}
