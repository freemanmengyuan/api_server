package base

import "fmt"

/**
 * 打印变量的类型和大小
*/
/*func main() {
	var a int = 89
	b := 95
	fmt.Println("value of a is", a, "and b is", b)
	fmt.Printf("type of a is %T, size of a is %d", a, unsafe.Sizeof(a)) // a 的类型和大小
	fmt.Printf("\ntype of b is %T, size of b is %d", b, unsafe.Sizeof(b)) // b 的类型和大小
}
*/

/**
 * float32
*/
/*func main() {
	a, b := 5.67, 8.97
	fmt.Printf("type of a %T b %T\n", a, b)
	sum := a + b
	diff := a - b
	fmt.Println("sum", sum, "diff", diff)

	no1, no2 := 56, 89
	fmt.Println("sum", no1+no2, "diff", no1-no2)
}*/

/**
 * complex
*/
/*func main() {
	c1 := complex(5, 7)
	c2 := 8 + 27i
	cadd := c1 + c2
	fmt.Println("sum:", cadd)
	cmul := c1 * c2
	fmt.Println("product:", cmul)
}*/

/**
 * string
*/
/*func main() {
	first := "Naveen"
	last := "Ramanathan"
	name := first +" "+ last
	fmt.Println("My name is",name)
}*/

/**
 * 数据类型转换
*/
func main() {
	i := 55      //int
	j := 67.8    //float64
	sum := i + int(j) //不同类型的数据不允许进行运算 这里需要强制转换
	var c int
	c = int(j) //赋值也是一样
	fmt.Println(sum, c)
}


