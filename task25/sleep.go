package main

import (
	"context"
	"fmt"
	"time"
)

const sec = 2

func sleep1(d time.Duration) {
	<-time.After(d)
}

func sleep2(d time.Duration) {
	start := time.Now()
	for {
		if time.Since(start) >= d {
			return
		}
	}
}

func sleep3(d time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), d)
	defer cancel()
	<-ctx.Done()
}

func sleep4(d time.Duration) {
	t := time.NewTimer(d)
	defer t.Stop()
	<-t.C
}

func main() {
	c := 1

	fmt.Printf("(%d) You'll see me in %d seconds\n", c, sec)
	sleep1(sec * time.Second)
	fmt.Printf("(%d) I missed you\n\n", c)
	c++

	fmt.Printf("(%d) You'll see me in %d seconds\n", c, sec)
	sleep2(sec * time.Second)
	fmt.Printf("(%d) I missed you\n\n", c)
	c++

	fmt.Printf("(%d) You'll see me in %d seconds\n", c, sec)
	sleep3(sec * time.Second)
	fmt.Printf("(%d) I missed you\n\n", c)
	c++

	fmt.Printf("(%d) You'll see me in %d seconds\n", c, sec)
	sleep4(sec * time.Second)
	fmt.Printf("(%d) I missed you\n", c)
}
