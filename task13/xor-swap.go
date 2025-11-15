// Program demonstrating XOR swap algorithm.
package main

import "fmt"

func main() {
	// User inputs a and b.
	var a, b int
	fmt.Print("Enter a: ")
	fmt.Scan(&a)
	fmt.Print("Enter b: ")
	fmt.Scan(&b)
	fmt.Println()

	// XOR swap algorithm:
	// 1) a = a XOR b
	// 2) b = a XOR b XOR b = a
	// 3) a = a XOR b XOR a = b
	a ^= b
	b ^= a
	a ^= b

	// Algorithm result output.
	fmt.Printf("Result: a = %d, b = %d", a, b)
	fmt.Println()
}
