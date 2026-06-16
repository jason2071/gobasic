package main

import "fmt"

func mostFrequent(nums []int) int {
	counts := make(map[int]int)

	for _, n := range nums {
		counts[n]++
	}

	best, maxCount := 0, 0
	for num, count := range counts {

		if count > maxCount {
			best = num
			maxCount = count
		}

	}

	return best
}

func main() {
	r := mostFrequent([]int{1, 3, 3, 3, 2, 1})
	fmt.Println(r)
}
