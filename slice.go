package main

import "fmt"

//切片不是值类型的 相当于数组的视图
func sliceInit() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s := arr[2:6] //半开半闭区间
	fmt.Println(s)

	fmt.Println(arr[2:6])
	fmt.Println(arr[:6])
	fmt.Println(arr[2:])
	fmt.Println(arr[:])
}

func sliceUpdate(s []int) {
	s[0] = 100
}

func main() {
	fmt.Println("hello")
	//初始化slice
	sliceInit()
	//slice相当于数组的视图 如果更该视图的值 相对应的数组的值也会更改
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s := arr[2:]
	sliceUpdate(s)
	fmt.Println(s)
	//更改数组的值
	sliceUpdate(arr[:])
	fmt.Println(arr)
	//重复的slice reslice
	ss := s[2:]
	fmt.Println(ss)
	//slice的扩展 slice可以向后扩展 不可以向前扩展
	//s[i]不可以超越len(s) 向后扩展不可以超越底层数组的cap(s)
	s1 := arr[2:6]
	s2 := s1[3:5]
	fmt.Println(s1)
	fmt.Println(s2)

}
