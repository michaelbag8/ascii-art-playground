package main

import (
	"unicode"
)


//1.
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

//2.
func CountVisibleChar(str []string) int{
	count := 0
	for _, word := range str{
		for _, ch := range word{
			if ch != ' '{
				count++
			}
		}
	}
	return count
	
}
