package main

import (
	"fmt"
	"unicode/utf8"
)

func isAnagram(a, b string) bool {
	if utf8.RuneCountInString(a) != utf8.RuneCountInString(b) {
		return false
	}

	counts := make(map[rune]int)
	for _, r := range a {
		counts[r]++
	}

	for _, r := range b {
		counts[r]--
		if counts[r] < 0 {
			return false
		}
	}

	return true
}

func main() {
	// fmt.Println(isAnagram("listen", "silent")) // true
	fmt.Println(isAnagram("hello", "world")) // false
}
