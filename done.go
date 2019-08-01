package main

import (
	"fmt"
	"sync"
)

//使用channel等待任务结束  通知外面任务完成
/*func dowork(id int, c chan int, done chan bool) {
	for a := range c {
		fmt.Printf("worker %d received %c \n", id, a)
		done <- true
	}
}

type work struct {
	in   chan int
	done chan bool
}

//返回channel
func createWork(id int) work {
	w := work{
		in:   make(chan int),
		done: make(chan bool),
	}
	go dowork(id, w.in, w.done)
	return w
}

func channelDemo() {
	var workers [10]work
	//done := make(chan bool)
	for i := 0; i < 10; i++ {
		workers[i] = createWork(i)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for _, worker := range workers {
		<-worker.done
	}
	//time.Sleep(time.Millisecond)
}*/

//使用系统提供的wg等待任务结束
func dowork(id int, c chan int, wg *sync.WaitGroup) {
	for a := range c {
		fmt.Printf("worker %d received %c \n", id, a)
		wg.Done()
	}
}

type work struct {
	in chan int
	wg *sync.WaitGroup
}

//返回channel
func createWork(id int, wg *sync.WaitGroup) work {
	w := work{
		in: make(chan int),
		wg: wg,
	}
	go dowork(id, w.in, w.wg)
	return w
}

func channelDemo() {
	var wg sync.WaitGroup
	var workers [10]work
	//done := make(chan bool)
	for i := 0; i < 10; i++ {
		workers[i] = createWork(i, &wg)
	}

	wg.Add(10)
	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	wg.Wait()
}

func main() {
	channelDemo()
}
