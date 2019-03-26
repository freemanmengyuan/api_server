package main

import "fmt"

func main() {
	//初始化map
	m := map[string]string{
		"name": "mengyuan",
		"age":  "16",
		"sex":  "1",
	}
	fmt.Println(m)
	m2 := make(map[string]string)
	fmt.Println(m2)
	var m3 map[string]string
	fmt.Println(m3)
}
