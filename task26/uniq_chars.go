package main

import (
	"fmt"
	"math/rand/v2"
	"unicode"
)

const (
	letters      = "абвгдеёжзийклмнопрстуфхцчшщъыьэюяАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ"
	stringLength = 7
)

type RuneSet map[rune]struct{}

// Add adds a rune to RuneSet object.
func (ss RuneSet) Add(char rune) {
	ss[char] = struct{}{}
}

func main() {
	// String init and print
	lettersRune := []rune(letters)
	strRune := make([]rune, stringLength)
	fmt.Print("String: ")
	for i := range strRune {
		strRune[i] = lettersRune[rand.IntN(len(lettersRune))]
		fmt.Print(string(strRune[i]))
	}
	fmt.Println()

	fmt.Println("Let's check if there are only uniq chars")
	if r, ok := AreCharsUniq(string(strRune)); !ok {
		fmt.Println("Letter", string(r), "is a duplicate!")
	} else {
		fmt.Println("All characters are unique!")
	}
}

// AreCharsUniq checks if there any duplicates in rune slice.
func AreCharsUniq(s string) (rune, bool) {
	seen := make(RuneSet, len(s))
	for _, r := range s {
		lowerRune := unicode.ToLower(r)
		if _, ok := seen[lowerRune]; ok {
			return r, false
		}
		seen.Add(lowerRune)
	}
	return 0, true
}
