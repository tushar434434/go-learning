package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("example.txt")

	if err != nil {
		panic(err)
	}
	fileInfo, err := f.Stat()

	if err != nil {
		panic(err)
	}
	fmt.Println("file name :", fileInfo.Name())
	fmt.Println("file or folder :", fileInfo.IsDir())
	fmt.Println("file size :", fileInfo.Size())
	fmt.Println("file permission :", fileInfo.Mode())
	fmt.Println("file modified at:", fileInfo.ModTime())

}
