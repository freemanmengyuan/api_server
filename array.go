package main

import "fmt"

func arrayInit() {
	var arr1 [5]int
	fmt.Println(arr1)
	arr2 := [5]int{1, 2, 3, 4, 5}
	arr3 := [...]int{1, 3, 5}
	var arr4 [4][5]int
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	var arr5 = [3]int{1, 2, 3}
	fmt.Println(arr5)
}

func arrayPrint() {
	arr2 := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}
	//比较美观 意义明确
	for i, v := range arr2 {
		fmt.Println(i, v)
	}
	for _, v := range arr2 {
		fmt.Println(v)
	}
}

func arrayPrintTest(arr [5]int) { //传参类型必须一致

	for i, v := range arr {
		fmt.Println(i, v)
	}
	arr[2] = 100
}

func arrayPrintPoint(arr *[5]int) {

	for i, v := range arr {
		fmt.Println(i, v)
	}
	arr[2] = 100
}

func main() {
	fmt.Println("hello")
	//初始化
	arrayInit()
	//遍历
	//arrayPrint()
	//数组是值类型 作为函数参数 是复制
	arr2 := [5]int{1, 2, 3, 4, 5}
	arrayPrintTest(arr2)
	fmt.Println(arr2)

	//使用指针 更改数组值
	arrayPrintPoint(&arr2)
	fmt.Println(arr2)
}
