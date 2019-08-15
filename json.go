package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	//数组
	x := []int{1, 2, 3, 4, 5}
	s, err := json.Marshal(x)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)

	//字典
	m := make(map[string]string)
	m["name"] = "mengyuan"
	m["age"] = "12"
	s1, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", s1)

}
