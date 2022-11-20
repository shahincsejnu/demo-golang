// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	var cnt int // cnt = 0

// 	var mu sync.Mutex

// 	go func() {
// 		mu.Lock()
// 		cnt++
// 		mu.Unlock()
// 	}()

// 	mu.Lock()
// 	if cnt == 0 {
// 		fmt.Println("Zero!")
// 	} else {
// 		fmt.Println("Not zero!")
// 	}
// 	mu.Unlock()
// }
