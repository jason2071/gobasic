package main

import (
	"fmt"
	"strings"
	_ "unicode"
)

func countVowels(s string) int {
	counts := 0
	for _, c := range s {
		lowerText := strings.ToLower(string(c))
		switch lowerText {
		case "a", "e", "i", "o", "u":
			counts++
		}

		// if strings.ContainsRune("aeiou", unicode.ToLower(c)) {
		// 	counts++
		// }
	}
	return counts
}

func main() {
	fmt.Println(countVowels("hello")) // 2
}
