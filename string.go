package main

import (
	"fmt"
)

func main() {
	s := "一蓑烟雨任平生" //UTF-8
	fmt.Printf("%X\n", []byte(s))

	for _, value := range []byte(s) {
		fmt.Printf("%X ", value)
	}
	fmt.Println()

	for i, ch := range s { //ch is a rune
		fmt.Printf("(%d  %X)", i, ch)
	}
	fmt.Println()

	//fmt.Println("dfdfd", utf8.DecodeLastRuneInString(s))

	for pos, value := range []rune(s) {
		fmt.Printf("(%d  %c)", pos, value)
	}
}
