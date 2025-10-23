// Our pipeline scheme (UNIX style):
// chGen | chSqr | stdout.
package main

import (
	"fmt"
	"math/rand"
)

const sliceLength = 15

// NumsGen writes random integers to rcvCh and then closes it.
// It generates sliceLength numbers in range [0, 100).
func NumsGen(rcvCh chan<- int) {
	for range sliceLength {
		rcvCh <- rand.Intn(100)
	}

	close(rcvCh)
}

// NumsSquare reads integers from sendCh, writes each number and its square to rcvCh,
// and then closes rcvCh.
func NumsSquare(sendCh <-chan int, rcvCh chan<- int) {
	for x := range sendCh {
		rcvCh <- x
		rcvCh <- x * x
	}

	close(rcvCh)
}

// NumsOutput prints to stdout the pairs (num, num^2) received from ansCh.
func NumsOutput(ansCh <-chan int) {
	for {
		num, ok := <-ansCh
		if !ok {
			break
		}

		numSqr, ok := <-ansCh
		if !ok {
			break
		}

		fmt.Printf("Num is %d, its square is %d\n", num, numSqr)
	}
}

func main() {
	chGen := make(chan int, 1)
	go NumsGen(chGen)

	chSqr := make(chan int, 2)
	go NumsSquare(chGen, chSqr)

	NumsOutput(chSqr)
}
