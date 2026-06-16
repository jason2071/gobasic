package main

import "fmt"

func factorial(number int) int {
	result := 1
	for i := 2; i <= number; i++ {
		result *= i
	}
	return result
}

func main() {
	fmt.Println(factorial(5)) // 120
	fmt.Println(factorial(4))
	fmt.Println(factorial(0)) // 1 (นิยาม: 0! = 1)
}
