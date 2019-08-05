package main

import (
	"fmt"
	"math/rand"
	"time"
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
func generator() chan int {
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
}
