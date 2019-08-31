package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func startHttpServer() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello, world"))
	})
	fmt.Print("server start....")
	http.ListenAndServe("127.0.0.1:8088", nil)
}

func sendGetRequest(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}

func sendPostRequest(url string, content string) string {
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(content))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}

func main() {

	//开启一个http服务
	//startHttpServer()
	//s := sendGetRequest("http://www.mengyuancc.cn")
	s := sendPostRequest("http://www.mengyuancc.cn", "id=3")
	fmt.Print(s)

}
