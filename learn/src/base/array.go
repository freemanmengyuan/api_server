package base

import "fmt"

func arrayInit() {
	// 完整的定义方式
	var arr0 [5]int
	var arr1 = [6]int{1, 3, 5, 7, 9, 0}
	fmt.Println(arr0)
	fmt.Println(arr1)
	// 短格式
	arr2 := [5]int{1, 2, 3, 4, 5}
	// 定义slice
	arr3 := [...]int{1, 3, 5}
	var arr4 [4][5]int
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	var arr5 = [3]int{1, 2, 3}
	fmt.Println(arr5)
}

func arrayForeach() {
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

func changearray(arr [5]int) { //传参类型必须一致
	for i, v := range arr {
		fmt.Println(i, v)
	}
	arr[2] = 100
}

func changeArrayByPoint(arr *[5]int) {

	for i, v := range arr {
		fmt.Println(i, v)
	}
	arr[2] = 100
}

func ArrayParamTest() {

	//数组是值类型 作为函数参数 是复制
	arr2 := [5]int{1, 2, 3, 4, 5}
	/*changearray(arr2)
	fmt.Println(arr2)*/

	//使用指针 更改数组值
	changeArrayByPoint(&arr2)
	fmt.Println(arr2)
}
