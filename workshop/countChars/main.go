package main

import (
	"errors"
	"fmt"
)

func main() {
	counts, err := countChars("hello")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(counts)
}

func countChars(s string) (map[rune]int, error) {
	if s == "" {
		return nil, errors.New("string is empty")
	}
	counts := make(map[rune]int)
	for _, c := range s {
		counts[c]++
	}
	return counts, nil
}
