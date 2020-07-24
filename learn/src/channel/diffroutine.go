package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)
	//协程1
	go func() {
		for i := 0; i < 100; i++ {
			if i == 10 {
				<-ch
			}
			fmt.Println(i)
		}
	}()

	//协程2
	go func() {
		for j := 100; j < 200; j++ {
			fmt.Println(j)
		}
	}()
	ch <- 1

	time.Sleep(time.Millisecond)
}
