// Program demonstrating incrementing in a competitive environment by using atomics.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const (
	goroutinesNum = 100
)

type Count struct {
	counter int32
}

func main() {
	count := &Count{}

	wg := sync.WaitGroup{}
	wg.Add(goroutinesNum)
	for range goroutinesNum {
		go func() {
			defer wg.Done()
			atomic.AddInt32(&count.counter, 1)
		}()
	}
	wg.Wait()

	fmt.Println("result:", count.counter)
}
