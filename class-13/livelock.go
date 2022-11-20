package main

// import (
// 	"bytes"
// 	"fmt"
// 	"sync"
// 	"sync/atomic"
// 	"time"
// )

// var mu sync.Mutex
// var left, right int32

// func wait() {
// 	mu.Lock()
// 	time.Sleep(1 * time.Millisecond)
// 	mu.Unlock()
// }

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(2)

// 	go walk(&wg, "Alice")
// 	go walk(&wg, "Bob")

// 	wg.Wait()
// }

// func walk(wg *sync.WaitGroup, name string) {
// 	var out bytes.Buffer
// 	defer func() {
// 		fmt.Println(out.String())
// 	}()

// 	defer wg.Done()

// 	fmt.Fprintf(&out, "%v is started walking:", name)

// 	for i := 0; i < 5; i++ {
// 		if tryLeft(&out) || tryRight(&out) {
// 			return
// 		}
// 	}

// 	fmt.Fprintf(&out, "\n%v gave up", name)
// }

// func tryLeft(out *bytes.Buffer) bool {
// 	return tryWalk("left", &left, out)
// }

// func tryRight(out *bytes.Buffer) bool {
// 	return tryWalk("right", &right, out)
// }

// func tryWalk(dirName string, count *int32, out *bytes.Buffer) bool {
// 	fmt.Fprintf(out, " %v", dirName)

// 	atomic.AddInt32(count, 1)

// 	wait()

// 	if atomic.LoadInt32(count) == 1 {
// 		fmt.Fprint(out, " yeah fine!")
// 		return true
// 	}

// 	wait()
// 	atomic.AddInt32(count, -1)
// 	return false
// }
