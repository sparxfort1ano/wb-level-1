// Program demonstrating string flipping.
package main

import (
	"fmt"
)

func main() {
	fmt.Print("input str: ")
	var str string
	fmt.Scan(&str)

	strRune := []rune(str)
	n := len(strRune)
	for i := 0; i < n/2; i++ {
		strRune[i], strRune[n-1-i] = strRune[n-1-i], strRune[i]
	}

	fmt.Println("flipped str:", string(strRune))
}
