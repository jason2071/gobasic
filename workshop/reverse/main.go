package main

import (
	"fmt"
	"strings"

	"github.com/rivo/uniseg"
)

func main() {
	fmt.Println(reverse("hello"))  // ได้ "olleh"
	fmt.Println(reverse("สวัสดี")) // ต้องไม่เพี้ยน (ระวังภาษาไทย/อักขระ Unicode)

	fmt.Println("----------------------")

	fmt.Println(reverseGraphemes("hello"))
	fmt.Println(reverseGraphemes("สวัสดี"))
}

func reverse(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func reverseGraphemes(s string) string {
	gr := uniseg.NewGraphemes(s)
	var cluster []string
	for gr.Next() {
		cluster = append(cluster, gr.Str())
	}

	for i, j := 0, len(cluster)-1; i < j; i, j = i+1, j-1 {
		cluster[i], cluster[j] = cluster[j], cluster[i]
	}

	var result strings.Builder
	for _, c := range cluster {
		result.WriteString(c)
	}

	return result.String()
}
