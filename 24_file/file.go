package main

import (
	"fmt"
	"os"
)

func main() {
	/*
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
		fmt.Println("file modified at:", fileInfo.ModTime())*/

	//read file

	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	// buf := make([]byte, 10)
	// d, err := f.Read(buf)
	// if err != nil {
	// 	panic(err)
	// }
	// for i := 0; i < len(buf); i++ {
	// 	println("data", d, string(buf[i]))
	// }

	// data, err := os.ReadFile("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(data))

	//read folder ======

	// dir, err := os.Open("../")
	// if err != nil {
	// 	panic(err)
	// }
	// defer dir.Close()
	// files, err := dir.Readdir(0)
	// if err != nil {
	// 	panic(err)
	// }

	// for _, file := range files {
	// 	fmt.Println(file.Name())
	// }

	//create a file
	// f, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	// f.WriteString("hi tushar")

	// bytes := []byte("hellooooooooo")
	// f.Write(bytes)

	// read and write to another file (streaming fashion)
	// sourceFile, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	// destFile, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer destFile.Close()

	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destFile)

	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}

	// 		break
	// 	}

	// 	e := writer.WriteByte(b)
	// 	if e != nil {
	// 		panic(e)
	// 	}
	// }
	// writer.Flush()
	// fmt.Println("written to new file succesfully")

	//delete a file
	sourceFile, err := os.Open("example2.txt")
	if err != nil {
		panic(err)
	}

	sourceFile.Close()

	err = os.Remove("example2.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("File deleted successfully")

}
