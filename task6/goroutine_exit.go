package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func naturalExit(wg *sync.WaitGroup) {
	defer wg.Done()

	var n int = 5
	for i := 0; i < n; i++ {
		time.Sleep(time.Second)
		log.Println("goroutine exiting in", n-i, "iterations")
	}
}

func conditionalExit(wg *sync.WaitGroup) {
	defer wg.Done()

	var x int = 0

	time.Sleep(time.Second)
	fmt.Println("When x equals 5, goroutine exits")

	for {

		log.Println("x equals", x)
		if x == 5 {
			return
		}
		time.Sleep(time.Second)

		x++
	}
}

func atomicExit(stopFlag *atomic.Bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for !stopFlag.Load() {
		time.Sleep(time.Second)
		log.Println("goroutine waiting for flag waving")
	}
}

func chanClosureExit(sendCh <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	ticker := time.NewTicker(time.Second)

	for {
		select {
		case <-sendCh:
			return
		case <-ticker.C:
			log.Println("goroutine waiting for channel closure")
		}
	}
}

func contextExit(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	ticker := time.NewTicker(time.Second)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			log.Println("goroutine waiting for context cancel")
		}
	}
}

func timeExit(timer *time.Timer, wg *sync.WaitGroup) {
	defer wg.Done()

	ticker := time.NewTicker(time.Second)

	for {
		select {
		case <-timer.C:
			return
		case <-ticker.C:
			log.Println("goroutine waiting for timeout")
		}
	}
}

func signalExit(sendCh <-chan os.Signal, wg *sync.WaitGroup) {
	defer wg.Done()

	ticker := time.NewTicker(2 * time.Second)

	for {
		select {
		case <-sendCh:
			return
		case <-ticker.C:
			log.Println("goroutine waiting for user's CTRL + C")
		}
	}
}

func runtimeExit(wg *sync.WaitGroup) {
	defer wg.Done()
	defer time.Sleep(time.Second) // imitation of cleanup work before exit
	defer fmt.Println("It also executes defer statements!")

	time.Sleep(time.Second)

	fmt.Println("This method immediately terminates the goroutine")
	time.Sleep(2 * time.Second)

	runtime.Goexit()

	fmt.Println("It's not executing!")
}

func main() {
	fmt.Print(`
Modes: 
1 -- Natural exit
2 -- Conditional exit
3 -- Conditional exit (atomic)
4 -- Exit after channel closure
5 -- Exit using context
6 -- Exit using time
7 -- Exit using SIGINT
8 -- Exit using runtime

`)

	var mode int = 0
	for mode < 1 || mode > 8 {
		fmt.Print("Enter a number from 1 to 8: ")
		fmt.Scan(&mode)
	}
	fmt.Println()

	var wg sync.WaitGroup

	wg.Add(1)

	switch mode {
	case 1:
		go naturalExit(&wg)

	case 2:
		go conditionalExit(&wg)

	case 3:
		var stopFlag atomic.Bool
		go atomicExit(&stopFlag, &wg)

		go func() { // auxiliary goroutine, waves the flag
			time.Sleep(6 * time.Second)
			stopFlag.Store(true)
		}()

	case 4:
		doneChan := make(chan struct{}) // empty struct channel -- perfect tool for signals
		go chanClosureExit(doneChan, &wg)

		go func() {
			time.Sleep(6 * time.Second)
			close(doneChan)
		}()

	case 5:
		ctx, cancel := context.WithCancel(context.Background())
		// Alternatives: context.WithDeadline(...), context.WithTimeout(...)

		go contextExit(ctx, &wg)

		go func() {
			time.Sleep(6 * time.Second)
			cancel()
		}()

	case 6:
		timer := time.NewTimer(6 * time.Second)
		// Alternatives: time.NewTicker(...), time.After(...)

		go timeExit(timer, &wg)

	case 7:
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt)

		go signalExit(sigChan, &wg)

	case 8:
		go runtimeExit(&wg)
	}

	log.Print("amount of goroutines: ", runtime.NumGoroutine(), "\n\n")
	wg.Wait()

	log.Print("goroutine successful exit\n\n")

	log.Println("amount of goroutines: 1")
}
