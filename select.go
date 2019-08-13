package main

import (
	"fmt"
)

/*func main() {
	//非阻塞
	var c1, c2 chan int
	for {
		select {
		case n := <-c1:
			fmt.Printf("receive from c1 %d", n)
		case n := <-c2:
			fmt.Printf("receive from c2 %d", n)
		default:
			fmt.Print("no value")
		}
	}
}*/
/*func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func main() {
	//非阻塞
	var c1, c2 = generator(), generator()

	for {
		select {
		case n := <-c1:
			fmt.Printf("receive from c1 %d\n", n)
		case n := <-c2:
			fmt.Printf("receive from c2 %d\n", n)
		}
	}
}*/
//main本身也是一个协程
/*func main() {
	ch := make(chan int)

	go func(ch chan int) {
		ch <- 1
	}(ch)

	//协程不会阻塞 直接向下执行到select
	time.Sleep(time.Millisecond)

	select {
	case <-ch:
		fmt.Println("hello")
	default:
		fmt.Println("default")
	}
}*/

func main() {
	ch := make(chan int)
	//timeout := make(chan int)

	go func(ch chan int) {
		ch <- 1
	}(ch)

	//协程不会阻塞 直接向下执行到select
	//time.Sleep(time.Millisecond)

	select {
	case <-ch:
		fmt.Println("hello")
	default:
		fmt.Println("default")
	}
}
