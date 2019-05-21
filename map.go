package main

import "fmt"

func initMap() {
	m := map[string]string{
		"name": "mengyuan",
		"age":  "16",
		"sex":  "1",
	}
	fmt.Println(m)
	m2 := make(map[string]string) //m2 = empty
	fmt.Println(m2)
	var m3 map[string]string    //m3 = nil
	fmt.Println(m3)
}

func foreachMap() {
	m := map[string]string{
		"name": "mengyuan",
		"age":  "16",
		"sex":  "1",
	}
	println("*****遍历*****")
	for k,v := range(m) {
		println(k, v)
	}
	//只拿key
	for k := range(m) {
		println(k)
	}
	//只拿value
	for _,v := range(m) {
		println(v)
	}
}

func getMapValue() {
	m := map[string]string{
		"name": "mengyuan",
		"age":  "16",
		"sex":  "1",
	}
	//println("*****获取元素*****")
	name := m["name"]
	fmt.Println(name)
	//判断健是否存在 如果键不存在 不报错 用ok来接进行判断
	if name, ok := m["class"]; ok {
		fmt.Println(name)
	} else {
		fmt.Println("key not exists")
	}
}

func delMapValue() {
	m := map[string]string{
		"name": "mengyuan",
		"age":  "16",
		"sex":  "1",
	}
	delete(m, "name")
	fmt.Println(m)
}

func main() {
	//初始化map
	initMap()
	//遍历map
	//foreachMap()
	//获取元素
	//getMapValue()
	//删除操作
	//delMapValue()
}
