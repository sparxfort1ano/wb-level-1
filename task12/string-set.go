// Program creating a set of words given in a slice.
package main

import (
	"fmt"
	"math/rand"
)

const (
	letters     = "abcdef"
	sliceLength = 20
)

// Set represents a set of strings.
type Set map[string]struct{}

// Add adds a string value to the set (no effect if already present).
func (s Set) Add(str string) {
	s[str] = struct{}{}
}

func main() {
	// Generating words of 2 letters randomly.
	words := make([]string, 0, sliceLength)
	for range sliceLength {
		words = append(words, string(letters[rand.Intn(len(letters))])+string(letters[rand.Intn(len(letters))]))
	}

	// Slice words output.
	fmt.Println("We got the following words:")
	for _, word := range words {
		fmt.Printf("%s ", word)
	}
	fmt.Println()
	fmt.Println()

	// Creating a set of words.
	wordsSet := make(Set, sliceLength)
	for _, word := range words {
		wordsSet.Add(word)
	}

	// Set wordsSet output.
	fmt.Println("Now let's remove the duplicates:")
	for word := range wordsSet {
		fmt.Printf("%s ", word)
	}
}
