// Set intersection program.
// Finds common elements of two integer slices.
// It works in O(n1 + n2), where n1 and n2 are slices sizes.
package main

import (
	"fmt"
	"log"
	"math/rand"
)

// Set represents a set of integers.
type Set map[int]struct{}

// Add adds an element to the set.
func (s Set) Add(val int) {
	s[val] = struct{}{}
}

// Contains checks if an element exists in the set.
func (s Set) Contains(val int) bool {
	_, exists := s[val]
	return exists
}

func main() {
	// User defines slices sizes.
	fmt.Print("Enter the size of first array: ")
	var n1 int
	if _, err := fmt.Scan(&n1); err != nil {
		log.Println(err)
		fmt.Println("...\nSomething went wrong")
		return
	}
	if n1 <= 0 {
		fmt.Println("Error: array length must be positive")
		return
	}

	fmt.Print("Enter the size of second array: ")
	var n2 int
	if _, err := fmt.Scan(&n2); err != nil {
		log.Println(err)
		fmt.Println("...\nSomething went wrong")
		return
	}
	if n2 <= 0 {
		fmt.Println("Error: array length must be positive")
		return
	}

	// Filling slices with random numbers.
	slice1 := make([]int, 0, n1)
	for range n1 {
		slice1 = append(slice1, rand.Intn(n1*3/2))
	}
	slice2 := make([]int, 0, n2)
	for range n2 {
		slice2 = append(slice2, rand.Intn(n2*3/2))
	}

	// We create unique containing unique elements of slice1.
	unique := make(Set, n1)
	for _, val := range slice1 {
		unique.Add(val)
	}

	// We check each element of slice2 to see if it is contained in the set (it is amortized O(1)).
	// If it does, we put the element in resulting set (result).
	result := make(Set, min(n1, n2))
	for _, val := range slice2 {
		if unique.Contains(val) {
			result.Add(val)
		}
	}

	// Output of source slices.
	fmt.Println("We found the intersection of the following arrays:")
	for _, val := range slice1 {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
	for _, val := range slice2 {
		fmt.Printf("%d ", val)
	}
	fmt.Println()

	// Special message, if we got empty intersection.
	if len(result) == 0 {
		fmt.Println("The interection is ∅.")
		return
	}

	// Intersection set output.
	fmt.Println("The intersection:")
	for val := range result {
		fmt.Printf("%d ", val)
	}
}
