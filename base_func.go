package main

import "fmt"

// 声明
func calculateBill(num int, price int) int {
	ret := num * price
	return ret
}

// 参数 和 返回值类型都是可选的
// 同类型的参数可以简写
func calculateBillSimple(num, price int) int {
	ret := num * price
	return ret
}

//多返回值

//命名返回值

//空白符

func main() {

	price, num := 4, 5                     //定义变量
	ret := calculateBillSimple(num, price) //调用
	fmt.Println(ret)
}
