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
func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

//命名返回值
func rectPropsName(length, width float64) (area float64, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return
}

func main() {

	price, num := 4, 5                     //定义变量
	ret := calculateBillSimple(num, price) //调用
	fmt.Println(ret)
	//空白符
	var a, _ = rectPropsName(2, 4)
	fmt.Printf("the area is %.2f \n", a)
}
