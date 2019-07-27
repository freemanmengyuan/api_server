package main

import (
	"fmt"
	"time"
)

/*func channelDemo() {

	//var c chan int //c=nil
	c := make(chan int)
	c <- 1
	c <- 2
	a := <-c
	fmt.Print(a)
}*/

//fatal error: all goroutines are asleep - deadlock!

//chan 作为参数传递
/*func worker(id int, c chan int) {

	for {
		a := <-c
		fmt.Printf("worker %d received %d", id, a)
		fmt.Println()
	}

}

func channelDemo() {
	c := make(chan int)
	go worker(0, c)
	c <- 1
	c <- 2
	time.Sleep(time.Minute)
	//fmt.Print(a)
}*/

//开多个work 进行分发
/*func worker(id int, c chan int) {
	for {
		a := <-c
		fmt.Printf("worker %d received %c", id, a)
		fmt.Println()
	}
}

func channelDemo() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go worker(0, channels[i])
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	time.Sleep(time.Millisecond)
	//fmt.Print(a)
}

func main() {
	channelDemo()
}*/

// channelbuffer channel的缓冲区
//如果没有接收方 加入buffer也不会	dead lock
/*func worker(id int, c chan int) {
	for {
		a := <-c
		fmt.Printf("worker %d received %c", id, a)
		fmt.Println()
	}
}

func channelbufferDemo() {

	c := make(chan int, 3)
	//go worker(0, c)

	c <- 'a'
	c <- 'b'
	c <- 'c'

	time.Sleep(time.Millisecond)
	//fmt.Print(a)
}

func main() {
	channelbufferDemo()
}
*/

//channelclose 必须在发送方close
//接收方
func worker(id int, c chan int) {
	//会打印空串
	/*for {
		a := <-c
		fmt.Printf("worker %d received %c", id, a)
		fmt.Println()
	}*/
	for {
		a, ok := <-c
		if ok {
			fmt.Printf("worker %d received %c", id, a)
			fmt.Println()
		}
	}
}

func channelcloseDemo() {

	c := make(chan int, 3)
	go worker(0, c)

	//发送方
	c <- 'a'
	c <- 'b'
	c <- 'c'
	close(c)

	time.Sleep(time.Millisecond)
	//fmt.Print(a)
}

func main() {
	channelcloseDemo()
}
