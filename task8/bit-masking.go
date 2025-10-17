package main

import "fmt"

// setBit sets the i-th bit (1-based from right) of mask to either 0 or 1.
//
// It works as follows:
//   - (mask & ^(1 << i)) clears the target bit unconditionally.
//   - (-bit & (1 << i)) produces either 0 or a mask with the target bit set.
//   - The final OR combines both operations.
//
// Mathematically: bit ⊂ {0,1} ⇒ -bit ⊂ {0, 0xFF...} => (-bit & (1 << i)) ⊂ {0, (1<<i)}.
func setBit(mask, i, bit int64) int64 {
	i--
	return (mask & ^(1 << i)) | (-bit & (1 << i))
}

func main() {
	var mask int64
	fmt.Print("Enter a number: ")
	fmt.Scan(&mask)

	var bit int64
	for {
		fmt.Print("Enter either 0 or 1: ")
		fmt.Scan(&bit)
		if bit == 0 || bit == 1 {
			break
		}
	}

	var pos int64
	for {
		fmt.Print("Enter a bit position (1-64): ")
		fmt.Scan(&pos)

		if pos >= 1 && pos <= 64 {
			break
		}
	}

	result := setBit(mask, pos, bit)
	fmt.Printf("The result is %d", result)
}
