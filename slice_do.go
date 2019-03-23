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

	fmt.Println("s3 = ", s3)
	fmt.Println("s4 = ", s4)
	fmt.Println("s5 = ", s5)
	fmt.Println("arr = ", arr)
}
func main() {
	//向slice中添加元素
	//如果元素超过cap，系统会重新分配更大的底层数组
	//由于是值传递的关系，必须接受append的返回值
	addSlice()
}
