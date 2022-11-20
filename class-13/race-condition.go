package main

// import (
// 	"fmt"
// 	"sync"
// )

// type Value struct {
// 	mu    sync.Mutex
// 	value int
// }

// func main() {
// 	var v1, v2 Value

// 	go printSum(&v1, &v2)

// 	go printSum(&v2, &v1)
// }

// func printSum(v1, v2 *Value) {
// 	v1.mu.Lock()
// 	v2.mu.Lock()

// 	fmt.Println(v1.value + v2.value)

// 	v1.mu.Unlock()
// 	v2.mu.Unlock()
// }
