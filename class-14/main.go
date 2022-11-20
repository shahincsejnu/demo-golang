package main

import (
	"fmt"
)

func main() {

	//runtime.GOMAXPROCS()

	ch1 := make(chan int)
	ch2 := make(chan int)

	go demo2(ch1)
	go demo1(ch2)

	for i := 0; i < 2; i++ {
		select {
		case val1 := <-ch1:
			fmt.Println(val1)
		case val2 := <-ch2:
			fmt.Println(val2)
		}
	}
}

func demo1(ch chan int) {
	ch <- 10
}

func demo2(ch chan int) {
	ch <- 20
}
