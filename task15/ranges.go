// Program demonstrating how to assign slices correctly
package main

import (
	"fmt"
	"math/rand"
)

const letters = "abcdefghijklmnopqrstuvwxyz"

// createHugeString generates a random string of length n filled with random Latin letters
func createString(n int) string {
	lettersStr := make([]byte, n)
	for i := range n {
		lettersStr[i] = letters[rand.Intn(len(letters))]
	}

	return string(lettersStr)
}

var justString string

// generateStringSlice extracts a x bytes string slice of a randomly generated string of length n safely
func generateStringSlice(x, n int) {
	v := createString(n)
	if len(v) >= x {
		justString = v[:x]
	}
}

func main() {
	x := rand.Intn(1025)
	n := rand.Intn(1025)
	fmt.Printf("x = %d, n = %d.\n", x, n)

	generateStringSlice(x, n)

	if len(justString) == 0 {
		fmt.Printf("If the length of v (%d) had not been checked to be greater than or equal to %d, a panic would have occurred.\n", n, x)
	}
	fmt.Printf("The length of justString is %d.\n", len(justString))
}
