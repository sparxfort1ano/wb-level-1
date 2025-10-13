package main

import (
	"log"
	"sync"
)

const N = 20

// SafeMap provides concurrent-safe access to standart map using RWMutex.
type SafeMap struct {
	m    map[int]int
	rwMu sync.RWMutex
}

// mapWriter concurrently writes to SafeMap and sync.Map
func mapWriter(safeMap *SafeMap, syncMap *sync.Map, wg *sync.WaitGroup) {
	defer wg.Done()

	var localWg sync.WaitGroup

	for i := 0; i < N; i++ {
		j := i

		localWg.Add(2)

		go func(j int) {
			defer localWg.Done()

			safeMap.rwMu.Lock()
			safeMap.m[j] = j * j
			safeMap.rwMu.Unlock()
		}(j)

		go func() {
			defer localWg.Done()

			syncMap.Store(j, j*j)
		}()
	}

	localWg.Wait()
}

// mapReader concurrently reads from SafeMap and sync.Map
func mapReader(safeMap *SafeMap, syncMap *sync.Map, wg *sync.WaitGroup) {
	defer wg.Done()

	var localWg sync.WaitGroup

	for i := 0; i < N; i++ {
		j := i

		localWg.Add(2)

		go func(j int) {
			defer localWg.Done()

			safeMap.rwMu.RLock()
			if val, ok := safeMap.m[j]; ok {
				log.Println(val, "(by SafeMap)")
			} else {
				log.Println("At", j, "iteration reader for SafeMap was faster")
			}
			safeMap.rwMu.RUnlock()
		}(j)

		go func() {
			defer localWg.Done()

			if val, ok := syncMap.Load(j); ok {
				log.Println(val, "(by sync.Map)")
			} else {
				log.Println("At", j, "iteration reader for sync.Map was faster")
			}
		}()
	}

	localWg.Wait()
}

func main() {
	safeMap := SafeMap{m: make(map[int]int, N)}
	var syncMap sync.Map

	var wg sync.WaitGroup
	wg.Add(2)

	go mapWriter(&safeMap, &syncMap, &wg)
	go mapReader(&safeMap, &syncMap, &wg)

	wg.Wait()

	log.Println("All goroutines finished successfully.")
}
