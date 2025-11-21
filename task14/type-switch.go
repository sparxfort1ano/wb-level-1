// Program demonstrating variable type defining at runtime using Type switch
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Random interface{} type variable initialization.
	var x interface{}
	switch rand.Intn(4) {
	case 0:
		x = make(chan float64, 1)
	case 1:
		x = 1
	case 2:
		x = "wb"
	case 3:
		x = true
	}

	// Type switch construction.
	switch x := x.(type) {
	case chan float64:
		fmt.Printf("Type: %T, Capacity: %d.", x, cap(x))
	case int:
		fmt.Printf("Type: %T, Value: %d.", x, x)
	case string:
		fmt.Printf("Type: %T, Value: %s.", x, x)
	case bool:
		fmt.Printf("Type: %T, Value: %t.", x, x)
	}
	fmt.Println()
}
