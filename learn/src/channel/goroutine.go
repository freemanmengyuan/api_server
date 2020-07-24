package main

import (
	"fmt"
	"time"
)

func main() {

	//抢占式的多任务处理
	for i := 1; i <= 10; i++ {
		go func(i int) {
			//for {
			fmt.Println(i)
			//}
		}(i)
	}
	time.Sleep(time.Millisecond)

	//非抢占式的多任务处理
	/*var a [100]int
	for i := 0; i < 100; i++ {
		go func(i int) {
			for {
				a[i]++
				runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Print(a)*/

}
