package main

import "fmt"

func unique(nums []int) []int {
	out := make([]int, 0, len(nums))
	seen := make(map[int]bool)

	for _, n := range nums {
		if !seen[n] {
			out = append(out, n)
			seen[n] = true
		}
	}
	return out
}

func main() {
	r := unique([]int{1, 2, 2, 3, 1}) // [1 2 3]
	fmt.Println(r)
}
