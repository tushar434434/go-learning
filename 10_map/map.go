package main

import (
	"fmt"
	"maps"
)

//  map -> key value pair

func main() {
	//creating map
	m := make(map[string]string)
	m["name"] = "tushar"
	m["age"] = "22"
	fmt.Println(m)
	fmt.Println(m["name"])
	fmt.Println(m["age"])
	fmt.Println(m["address"]) //if key is not present it will return zero value of that type
	n := map[string]int{"name": 22}
	fmt.Println(n)
	fmt.Println(n["len"])
	delete(n, "name") //deleting key value pair
	fmt.Println(n)
	clear(m) //clearing map
	fmt.Println(m)
	l := map[string]string{"name": "tushar", "age": "22"}
	fmt.Println(l)
	v, ok := l["name"]
	fmt.Println(v)
	if ok {
		fmt.Println("key present")
	} else {
		fmt.Println("key not present")
	}
	//comparing map
	m1 := map[string]string{"name": "tushar", "age": "22"}
	m2 := map[string]string{"name": "tushar", "age": "22"}
	fmt.Println(maps.Equal(m1, m2)) //map cannot be compared

}
