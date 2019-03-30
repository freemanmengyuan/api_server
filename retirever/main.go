package main

import (
	"fmt"
	"go/retirever/real"
)
//定义接口
type Retriever interface {
	Get(url string) string
}

//使用
func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}
func main() {
	//var r Retriever
	//r = mock.Retriever{"this is fake imooc"}
	//fmt.Println(download(r))

	fmt.Println(download(real.Retriever{}))


}
