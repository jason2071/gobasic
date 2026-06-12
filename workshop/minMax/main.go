package main

import (
	"errors"
	"fmt"
)

func main() {
	min, max, err := minMax([]int{}) // ได้ 1, 9
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(min, max)

}

func minMax(numbers []int) (int, int, error) {
	if len(numbers) == 0 {
		return 0, 0, errors.New("numbers is empty")
	}

	min, max := numbers[0], numbers[0]

	for _, v := range numbers {

		if v < min {
			min = v
		}

		if v > max {
			max = v
		}
	}

	return min, max, nil
}
