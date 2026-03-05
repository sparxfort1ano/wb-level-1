// Program demonstrating inline quicksort realization.
package main

import (
	"fmt"
	"math/rand/v2"
)

const (
	maxSliceSize = 30
	maxIntValue  = 30
)

// quickSort returns a sorted copy of slice.
func quickSort(slc []int) []int {
	if len(slc) <= 1 {
		return slc
	}

	slcRes := make([]int, len(slc))
	n := copy(slcRes, slc)

	left, right := 0, n-1
	quickSortAlgorithm(slcRes, left, right)

	return slcRes
}

// quickSortAlgorithm implements recursive in-place algorithm realization.
func quickSortAlgorithm(slc []int, left, right int) {
	if right <= left {
		return
	}

	barrierElem := slc[left+(right-left)/2]
	i, j := left, right

	for i <= j {
		for slc[i] < barrierElem {
			i++
		}
		for slc[j] > barrierElem {
			j--
		}
		if i <= j {
			swap(slc, i, j)
			i++
			j--
		}
	}

	quickSortAlgorithm(slc, left, j)
	quickSortAlgorithm(slc, i, right)
}

// swap exchanges two elements in slice.
func swap(slc []int, i, j int) {
	slc[i], slc[j] = slc[j], slc[i]
}

func main() {
	fmt.Print("Slice length: ")
	n := rand.IntN(maxSliceSize)
	fmt.Println(n)

	slc := make([]int, n)
	fmt.Print("Slice: ")
	for i := range n {
		slc[i] = rand.IntN(maxIntValue)
		fmt.Printf("%d ", slc[i])
	}
	fmt.Println()

	slcSorted := quickSort(slc)
	fmt.Print("Slice sorted: ")
	for _, val := range slcSorted {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
}
