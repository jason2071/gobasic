package main

import "fmt"

func evens(nums []int) []int {
	result := []int{}
	for _, n := range nums {
		if n%2 == 0 {
			result = append(result, n)
		}
	}
	return result
}

func main() {

	r := evens([]int{1, 2, 3, 4, 5, 6})
	fmt.Println(r)

}
