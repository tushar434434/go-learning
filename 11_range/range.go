package main

import "fmt"

func main() {
	// for iteration over data structures
	nums := []int{1, 2, 3, 4, 5}
	for i, num := range nums {
		fmt.Println(i, num)
	}
	for i := 0; i < len(nums); i++ {
		fmt.Println(i, nums[i])
	}
	sum := 0
	for l, num := range nums {
		fmt.Println(l)
		sum += num
	}
	fmt.Println(sum)
	m := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range m {
		fmt.Println(k, v)

	}
	for i, c := range "golang" { //unicode code point rune
		fmt.Println(i, c)
		fmt.Println(i, string(c))
	}
}
