package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Value struct {
	mu    sync.Mutex
	value int
}

func main() {
	var v1, v2 Value

	wg.Add(2)
	go printSum(&v1, &v2)

	go printSum(&v2, &v1)

	wg.Wait()
}

func printSum(v1, v2 *Value) {
	v1.mu.Lock()

	time.Sleep(2 * time.Second)

	v2.mu.Lock()

	fmt.Println(v1.value + v2.value)

	v1.mu.Unlock()
	v2.mu.Unlock()

	wg.Done()
}
