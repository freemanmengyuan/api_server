package main

import (
	"fmt"
	"github.com/astaxie/goredis"
)

func main(){

	var client goredis.Client
	client.Addr = "144.34.144.200:6379"
	/*********************string*******************************/
	//写入数据
	/*err := client.Set("token", []byte("accersdere453445"))
	if err != nil {
		panic(err)
	}*/
	//获取数据
	res,err := client.Get("token")
	if err != nil {
		panic(err)
	}
	fmt.Print(string(res))

	/**********************hash*******************************/
	/*m := map[string]interface{}{
		"name" : "mengyuan",
		"age" : 16,
		"sex" : 1,
	}*/
	//写入数据
	/*err = client.Hmset("mtest", m)
	if err != nil {
		panic(err)
	}*/
	//读取数据
	res,err = client.Hget("mtest", "name")
	if err != nil {
		panic(err)
	}
	fmt.Print(string(res))
}
