// Program demonstrating binary search algorithm realization.
// Implement left and right boundary search for target index.
package main

import (
	"fmt"
	"math/rand/v2"
	"sort"
)

const (
	maxSliceSize = 20
	maxIntValue  = 20

	targetNotFound = -1
)

// leftBoundary implement lower boundary binary search.
func leftBoundary(slc []int, x int) int {
	left, right := -1, len(slc)

	for right-left > 1 {
		mid := left + (right-left)/2
		if x > slc[mid] {
			left = mid
		} else {
			right = mid
		}
	}

	if (right >= 0 && right < len(slc)) && slc[right] == x {
		return right
	}
	return targetNotFound
}

// rightBoundary implement upper boundary binary search.
func rightBoundary(slc []int, x int) int {
	left, right := -1, len(slc)

	for right-left > 1 {
		mid := left + (right-left)/2
		if x >= slc[mid] {
			left = mid
		} else {
			right = mid
		}
	}

	if (left >= 0 && left < len(slc)) && slc[left] == x {
		return left
	}
	return targetNotFound
}

func main() {
	n := rand.IntN(maxSliceSize)
	fmt.Printf("Slice length: %d\n", n)

	slc := make([]int, n)
	for i := range n {
		slc[i] = rand.IntN(maxIntValue)
	}
	sort.Ints(slc)
	fmt.Print("Slice (must be sorted): ")
	for _, elem := range slc {
		fmt.Printf("%d ", elem)
	}
	fmt.Println()

	x := rand.IntN(maxIntValue)
	fmt.Printf("Target: %d\n", x)

	res1 := leftBoundary(slc, x)
	fmt.Printf("Left boundary target index: %d\n", res1)

	res2 := rightBoundary(slc, x)
	fmt.Printf("Right boundary target index: %d\n", res2)
}
