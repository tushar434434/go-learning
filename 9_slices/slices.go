package main

import (
	"fmt"
	"slices"
)

// slices are dynamic arrray
func main() {
	//uninitiallized slice is nill value
	var nums []int
	fmt.Println(len(nums))
	fmt.Println(nums)
	var num = make([]int, 2)
	fmt.Println(num)
	fmt.Println(len(num))
	fmt.Println(cap(num))     //cap = capacity maximum size of slice
	var n = make([]int, 2, 5) //2 = length 5 = capacity
	fmt.Println(n)
	fmt.Println(len(n))
	fmt.Println(cap(n))
	n = append(n, 1)
	var n1 = make([]int, len(n)) //copying slice

	fmt.Println(n1)
	n = append(n, 2)
	n = append(n, 3)
	n = append(n, 4)
	n = append(n, 5)
	n = append(n, 6)
	fmt.Println(n)
	fmt.Println(len(n))
	fmt.Println(cap(n))

	//slic  eoperator
	var s = []int{1, 2, 3, 4, 5}
	fmt.Println(s)
	fmt.Println(s[1:3])
	fmt.Println(s[:3])
	fmt.Println(s[1:])

	//slice package
	var s1 = []int{1, 2, 3, 4, 5}
	var s2 = []int{6, 7, 8, 9, 10}
	fmt.Println(slices.Equal(s1, s2))
	//2d slice
	var twoD = [][]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(twoD)

}
