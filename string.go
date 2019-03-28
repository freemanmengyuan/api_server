package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//定义字符串
	s := "一蓑烟雨任平生" //UTF-8
	fmt.Printf("%X\n", []byte(s))

	//遍历字符串
	for _, value := range []byte(s) {
		fmt.Printf("%X ", value)
	}
	fmt.Println()

	for i, ch := range s { //ch is a rune
		fmt.Printf("(%d  %X)", i, ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d  %c)", i, ch)
	}
	fmt.Println()

	//获取字符串的字数
	count := len([]rune(s))
	println(count)

	fmt.Println(utf8.DecodeLastRuneInString(s))
}
