package main

import "fmt"

func printSlice[T any](items []T) { //generic function
	for _, item := range items {
		fmt.Println(item)
	}
}
func printS[T int | string](items []T) { // selective type constraint
	for _, item := range items {
		fmt.Println(item)
	}
}
func prin[T comparable](items []T) { //generic function comparable constraint
	for _, item := range items {
		fmt.Println(item)
	}
}

type stack[T any] struct {
	element []T
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	printSlice(slice)
	printS(slice)
	prin(slice)
	mystack := stack[int]{element: []int{1, 2, 3, 4, 5}}
	fmt.Println(mystack.element)
	mystack2 := stack[string]{element: []string{"a", "b", "c", "d", "e"}}
	fmt.Println(mystack2.element)
}
