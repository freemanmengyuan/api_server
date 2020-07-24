package tree

import "fmt"

//定义结构
type Node struct {
	Value int
	Left, Right *Node
}

//定义方法
//打印
func (node Node) Print() {  //和正常的写法没有区别
	fmt.Println(node.Value)
}

//更改 改进 加判断nil 值接受者 指针接受者
func (node *Node)SetValue(){
	if node == nil {
		fmt.Println("setting is error nil tree")
		return
	}
	node.Value = 0
}

//遍历 中序遍历的例子 先左后中后右
func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}