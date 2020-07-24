package base

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

// defer
func clear() {
	println("清理一些系统资源。。。。。。")
}

