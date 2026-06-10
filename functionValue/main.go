package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func adder() func(int) int {
	total := 0
	return func(i int) int {
		total += i
		return total
	}
}

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		result := a
		a, b = b, a+b
		return result
	}
}

func main() {
	// Example 1: ใช้ function literal (anonymous function)
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	// Example 2: ใช้ math.Pow
	result1 := compute(hypot)    // hypot(3, 4) = 5.0
	result2 := compute(math.Pow) // math.Pow(3, 4) = 81.0

	fmt.Println(result1) // 5
	fmt.Println(result2) // 81

	fmt.Println("----------------")

	add := adder()
	fmt.Println(add(1))
	fmt.Println(add(2))
	fmt.Println(add(3))
	fmt.Println(add(4))

	fmt.Println("----------------")

	f := fibonacci()
	for range 10 {
		fmt.Println(f())
	}
}
