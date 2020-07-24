package main

import (
	"learn/tree"
)

//包和封装
//使用大坨峰法命名 小驼峰代表private 大驼峰代表public
//一个目录代表一个包 目录名和包名可以不一致

func main() {
	//初始化
	var root tree.Node
	//root.Print()
	//return
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{4, nil, nil}
	root.Right = &tree.Node{5, nil, nil}
	root.Left.Right = new(tree.Node)

	root.Traverse()
}

//go语言的工作目录
//go build main.go
//go install hello
//go run main.go
