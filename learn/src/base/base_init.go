package base

import "fmt"

var (
	name = "mengyuan"
	age  = 12
)

//变量的定义
//只定义 不赋值
func variableZeroValue() {
	var a int
	var s string
	s = "hello"
	//fmt.Println(a, s)
	fmt.Printf("%d %s\n", a, s)
}

//定义并且赋值
func variableInitValue() {
	var a, b int = 3, 4
	var s string = "hello"
	fmt.Println(a, b, s)
}

//系统判别类型
func variableSystemValue() {
	var a, b, s = 3, 4, "world"
	fmt.Println(a, b, s)
}

//短类型
func variableShortValue() {
	a, b, s := 3, 4, "hahah"
	fmt.Println(a, b, s)
}

//第一个程序 helloworld
func main() {
	var(
		b = 2
		c = 2
	)
	fmt.Println(b, c)
	fmt.Println("hello world1")
	variableZeroValue()
	//variableInitValue()
	//variableSystemValue()
	//variableShortValue()
	//fmt.Println(name, age)

}
