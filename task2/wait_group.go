package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := [...]int{2, 4, 6, 8, 10}
	resSlice := make([]int, len(arr))

	var wg sync.WaitGroup
	wg.Add(len(arr))

	for i, val := range arr {
		i, val := i, val // shadowing
		go func() {
			defer wg.Done()
			resSlice[i] = val * val
		}()
	}

	wg.Wait()

	for _, val := range resSlice {
		fmt.Println(val) // вывод будет в строго последовательном порядке
	}
}
