package main

import (
	"fmt"
	"log"
	"math/rand/v2"
)

const sliceLength = 10

func main() {
	sl := make([]int, sliceLength)
	fmt.Print("Slice:    ")
	for i := range sl {
		sl[i] = rand.IntN(20)
		fmt.Printf(" %d", sl[i])
	}
	fmt.Println()

	toDeleteIndex := rand.IntN(len(sl))
	fmt.Printf("Let's delete %d element: %d\n",
		toDeleteIndex, sl[toDeleteIndex])
	if err := removeElemByIndex(&sl, toDeleteIndex); err != nil {
		log.Println(err.Error())
		return
	}

	fmt.Print("New slice:")
	for i := range sl {
		fmt.Printf(" %d", sl[i])
	}
	fmt.Println()
}

func removeElemByIndex(sl *[]int, i int) error {
	if i < 0 || i >= len(*sl) {
		return fmt.Errorf("index out of range: %d with slice length %d", i, len(*sl))
	}
	if len(*sl) == 0 {
		return fmt.Errorf("empty slice")
	}
	copy((*sl)[i:], (*sl)[i+1:])
	*sl = (*sl)[:len(*sl)-1]
	return nil
}
