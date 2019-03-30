package tree

import "fmt"

//定义结构
type Node struct {
	Value int
	Left, Right *Node
}

//打印
func (node Node) Print() {  //和正常的写法没有区别
	fmt.Println(node.Value)
}
//更改 改进 加判断nil
func (node *Node)SetValue(){
	if node == nil {
		fmt.Println("setting is error nil node")
		return
	}
	node.Value = 0
}

//遍历
func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}