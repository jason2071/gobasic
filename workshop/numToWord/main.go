package main

import "fmt"

func numToWord2(n int) string {
	words := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	if n < 1 || n > len(words) {
		return "unknown"
	}

	return words[n-1]
}

func numToWord(n int) string {
	words := []string{
		"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}
	if n < 0 || n > 9 {
		return "unknown"
	}
	return words[n]
}

func main() {
	fmt.Println(numToWord(3))  // "three"
	fmt.Println(numToWord(15)) // "unknown"
	fmt.Println(numToWord2(9))
}
