package main

import "fmt"

func sumDigits(n int) int {
	sum := 0

	// if n < 0 {
	// 	n = -n
	// }

	for n > 0 {
		sum += n % 10
		n = n / 10
	}

	return sum
}

func main() {
	fmt.Println(sumDigits(123))  // 1+2+3 = 6
	fmt.Println(sumDigits(4920)) // 4+9+2+0 = 15
	// fmt.Println(sunDigits(-123))
}
