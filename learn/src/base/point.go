package base

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func swapNew(a, b int) (int, int) {
	a, b = b, a
	return a, b
}

func PointDemo() {
	// 指针的典型例子 间接的更改变量的值 是指针存在的根本意义
	var a int = 2
	var pa *int = &a
	*pa = 3
	fmt.Println(a)

	// go只有值传递 没有引用传递
	// 交换变量的值
	a, b := 3, 4
	//swap(&a, &b)
	a, b = swapNew(a, b)
	fmt.Println(a, b)
}
