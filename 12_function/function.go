package main

import "fmt"

func add(a int, b int) int {
	return a + b
}
func getLanguages() (string, string, int) {
	return "golang", "python", 3
}

// func process(fn func(a int) int) {
// 	fn(1)
// }
func process() func(a int) int {
	return func(a int) int {
		return 2
	}
}

func main() {
	result := add(5, 6)
	fmt.Println(result)
	getLanguages()
	fmt.Println(getLanguages())
	lang1, lang2, lang3 := getLanguages()
	fmt.Println(lang1, lang2, lang3)
	// fn := func(a int) int {
	// 	return 2
	// }
	// process(fn)
	// fmt.Println(fn(2))
	fn := process()
	fmt.Println(fn(2))

}
