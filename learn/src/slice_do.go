package main

import "fmt"

func addSlice() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr)
	s1 := arr[2:6]
	s2 := s1[3:5]
	fmt.Println(s1)
	fmt.Println(s2)

	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)

	fmt.Println("s3 = ", s3, cap(s3))
	fmt.Println("s4 = ", s4, cap(s4))
	fmt.Println("s5 = ", s5, cap(s5))
	fmt.Println("arr = ", arr)
}

func initSlice() {
	var s []int
	for i := 1; i <= 100; i++ {
		fmt.Printf("len = %d cap = %d\n", len(s), cap(s))
		s = append(s, 2*i+1)
	}
	fmt.Println(s)
	s1 := []int{2, 4, 6, 8}
	s2 := make([]int, 10, 32)
	fmt.Printf("len = %d cap = %d v= %v\n", len(s1), cap(s1), s1)
	fmt.Printf("len = %d cap = %d v= %v\n", len(s2), cap(s2), s2)
	//拷贝操作
	copy(s2, s1)
	fmt.Printf("len = %d cap = %d v= %v\n", len(s2), cap(s2), s2)
	//删除操作
	s2 = append(s2[:3], s2[4:]...) //go语言中没有数组相加的概念 需要用append
	fmt.Printf("len = %d cap = %d v= %v\n", len(s2), cap(s2), s2)
	//pop操作
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Printf("tail = %d\n", tail)
	fmt.Printf("len = %d cap = %d v= %v\n", len(s2), cap(s2), s2)
}

func main() {
	//向slice中添加元素
	//如果元素超过cap，系统会重新分配更大的底层数组 超过后s4,s5view的将不会再是arr
	//由于是值传递的关系，必须接受append的返回值
	addSlice()
	//初始化slice
	initSlice()
}
