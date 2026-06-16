package main

import (
	"fmt"
	"strings"
)

func wordCountLoop(s string) int {
	counts := 0
	for _, words := range strings.Fields(s) {
		if len(words) > 0 {
			counts++
		}
	}

	return counts
}

func wordCount(s string) int {
	words := strings.Fields(s)
	return len(words)
}

func main() {
	fmt.Println(wordCount("the quick brown fox")) // ได้ 4
	fmt.Println(wordCount("  hello   world  "))   // ได้ 2 (ช่องว่างเกินไม่นับ)
	fmt.Println(wordCount(""))
}
