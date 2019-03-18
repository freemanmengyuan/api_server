package main

import "fmt"

func ifDemo() {

	// 简写的方式
	if a := 12; a%2 == 0 { // num 的作用域在if中
		fmt.Printf("this number is even %d\n", a)
	} else {
		fmt.Printf("this number is odd %d\n", a)
	}
}

func switchDemo() {
	//num := 5
	switch num := 2; num { // num 的作用域在switch
	case 1:
		fmt.Println("php")
	case 2:
		fmt.Println("python")
	case 3:
		fmt.Println("java")
	// case 3: case不允许出现重复项
	//	fmt.Println("js")
	case 4:
		fmt.Println("cpp")
	case 5:
		fmt.Println("golang")
	default: // default可以写在任何位置
		fmt.Println("default")
	}

	//无表达式的switch  相当于switch true
	nums := 2
	switch {
	case nums <= 2:
		fmt.Println("E")
	case nums <= 5:
		fmt.Println("D")
	case nums <= 6:
		fmt.Println("C")
	//case 3: case不允许出现重复项
	//	fmt.Println("js")
	case nums <= 8:
		fmt.Println("B")
	case nums <= 10:
		fmt.Println("A")
	default:
		fmt.Println("default")
	}
	fmt.Println(nums)

	// 包含多个表达式
	numbers := 5
	switch numbers {
	case 2, 3, 4, 5:
		fmt.Println("the number <= 5")
	case 6, 7, 8, 9, 10:
		fmt.Println("the number <= 5")
	}

	// Fallthrough 语句 在 Go 中，每执行完一个 case 后，会从 switch 语句中跳出来，不再做后续 case 的判断和执行。
	// 使用 fallthrough 语句可以在已经执行完成的 case 之后，把控制权转移到下一个 case 的执行代码中。
	number := 5
	switch number {
	case 2:
		fmt.Println("E")
		fallthrough
	case 5:
		fmt.Println("D")
		fallthrough
	case 6:
		fmt.Println("C")
		fallthrough
	//case 3: case不允许出现重复项
	//	fmt.Println("js")
	case 8:
		fmt.Println("B")
		fallthrough
	case 10:
		fmt.Println("A")
	}
}

func forDemo() {
	// for循环语句包括 初始条件 终止条件 post语句 三部分
	//num := 10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// continue 终止本次循环
	//num := 10
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	// break 跳出整个循环
	//num := 10
	for i := 1; i <= 10; i++ {
		if i > 7 {
			break
		}
		fmt.Println(i)
	}

	// 模仿while
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	//无限循环
	for {
		fmt.Println("hello")
	}
}
func main() {
	// 分支
	//ifDemo()
	// switch
	//switchDemo()
	// for循环
	forDemo()

}
