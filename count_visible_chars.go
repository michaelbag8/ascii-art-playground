package main

import (
	"unicode"
)

func CountVisibleChars(str []string) int {
	count := 0
	for _, w := range str {
		runes := []rune(w)
		for _, c := range runes {
			if unicode.IsLetter(c) || unicode.IsPunct(c) || unicode.IsDigit(c){
				count++
			}
		}
	}

	return count
}