
package base

import "fmt"

/**
 * 定义常量 指定类型
*/
/*func main() {
	const name = "Sam"
	const typename string = "hello"
	fmt.Printf("type %T value %v\n", name, name)
	fmt.Println(typename)
}*/

/**
 * 强类型语言赋值操作需要两边的类型一致
*/
/*func main() {
	const trueConst = true //类型不确定
	type myBool bool
	var defaultBool bool = trueConst // 允许 在赋值时确定类型
	var customBool myBool = trueConst // 允许
	//defaultBool = customBool // 不允许

	fmt.Println(defaultBool, customBool)
}*/


/**
 * 数字表达式
*/
func main() {

	const(
		name = "hell"
	)

	var a = 5.9/8
	fmt.Printf("a's type %T value %v name is %s",a, a, name)
}