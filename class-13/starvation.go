package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// var wg sync.WaitGroup
// var mu sync.Mutex

// const t = 1 * time.Second

// func main() {
// 	wg.Add(2)

// 	go inefficientWorker()

// 	go efficientWorker()

// 	wg.Wait()
// }

// func inefficientWorker() {
// 	defer wg.Done()

// 	var count int

// 	for begin := time.Now(); time.Since(begin) <= t; {
// 		mu.Lock()
// 		time.Sleep(3 * time.Nanosecond)
// 		mu.Unlock()

// 		count++
// 	}

// 	fmt.Println("inefficientWorker's loop count:", count)
// }

// func efficientWorker() {
// 	defer wg.Done()

// 	var count int

// 	for begin := time.Now(); time.Since(begin) <= t; {
// 		mu.Lock()
// 		time.Sleep(1 * time.Nanosecond)
// 		mu.Unlock()

// 		mu.Lock()
// 		time.Sleep(1 * time.Nanosecond)
// 		mu.Unlock()

// 		mu.Lock()
// 		time.Sleep(1 * time.Nanosecond)
// 		mu.Unlock()

// 		count++
// 	}

// 	fmt.Println("efficientWorker's loop count:", count)
// }
