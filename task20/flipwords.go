// Program demonstrating words flipping in a string (inline).
package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("input words in one line:")
	inputStr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	// symbols reverting
	runeStr := []rune(inputStr[:len(inputStr)-1])
	reverseString(runeStr)
	fmt.Println("flipped sequence of words:", string(runeStr))

	var start, end int
	for i, letter := range runeStr {
		// word scan and its reverting
		if unicode.IsSpace(letter) {
			if end-start != 0 {
				reverseString(runeStr[start:end])
			}
			start = i + 1
			end = i + 1
		} else {
			end++
		}
	}

	reverseString(runeStr[start:])

	fmt.Println("flipped sequence of words:", string(runeStr))
}

func reverseString(str []rune) {
	n := len(str)
	fmt.Println(string(str))
	for i := 0; i < n/2; i++ {
		str[i], str[n-1-i] = str[n-1-i], str[i]
		fmt.Println(string(str))
	}
}
