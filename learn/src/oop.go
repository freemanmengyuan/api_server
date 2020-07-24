package main

import "fmt"

//go语言的面向对象 只有封装 没有继承和多态

//定义结构
type treeNode struct {
	value int
	left, right *treeNode
}

//初始化
func structInit() {
	type treeNode struct {
		value int
		left, right *treeNode
	}
	//初始化
	var root treeNode
	root = treeNode{value:3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.left.right = new(treeNode)

	fmt.Println(root)
	//结构数组 相当于slice
	node := []treeNode {
		{value:3},
		{},
		{6, nil, &root},
	}
	fmt.Println(node)
}
//初始化结构也可以使用函数来实现
func createTreeNode(value int) *treeNode {
	return &treeNode{value:value}  //返回指针 是局部的 但也可以使用
}

//打印
func (node treeNode) print() {  //和正常的写法没有区别
	fmt.Println(node.value)
}
//更改 改进 加判断nil
func (node *treeNode)setValue(){
	if node == nil {
		fmt.Println("setting is error nil tree")
		return
	}
	node.value = 0
}

//遍历
func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func main() {
	//结构的定义和初始化
	//structInit()
	//利用函数来初始化结构
	node := treeNode{3, createTreeNode(5), nil}
	fmt.Println(node)
	//值传递 引用传递
	//tree.print()
	node.left.print()
	//更改 值传递||引用传递
	node.setValue()
	node.print()
	//传递空指针
	var node1 *treeNode
	node1.setValue()
	//遍历
	//初始化
	var root treeNode
	root = treeNode{value:3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.left.right = new(treeNode)
	root.traverse()

}


