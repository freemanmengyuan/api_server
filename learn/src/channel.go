package main

import (
	"fmt"
)

//利用通信来共享内存 而不是利用共享内存来通信
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
func worker(id int, c chan int) {
	for {
		//sfmt.Printf("worker %d statr\n", id)
		a := <-c
		fmt.Printf("worker %d received %c", id, a)
		fmt.Println()
	}
}

func channelDemo() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go worker(i, channels[i])
	}

	for {
		for i := 0; i < 10; i++ {
			channels[i] <- 'a' + i
		}
	}

	//time.Sleep(time.Minute)
	//fmt.Print(a)
}

func main() {
	channelDemo()
}

//chan 作为返回值
/*func work(id int, c chan int) {
	for a := range c {
		fmt.Printf("worker %d received %c \n", id, a)
		//done <- true
		//done chan bool
	}

}
//返回channel
func createWork(id int) chan<- int {
	c := make(chan int)
	go work(id, c)
	return c
}

func channelDemo() {
	var channels [10]chan<- int
	//done := make(chan bool)
	for i := 0; i < 10; i++ {
		channels[i] = createWork(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	time.Sleep(time.Millisecond)
	//fmt.Print(a)
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
