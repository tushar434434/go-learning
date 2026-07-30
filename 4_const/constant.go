package main

import "fmt"

func main() {
	const name = "tushar"
	fmt.Println(name)

	const (
		port = 50
		host = "local host"
	)
	fmt.Println(port, host)

}
