package main

import "fmt"

func main() {
	sum := 0
	for i := range 10 {
		sum += i // sum = sum + i
	}
	fmt.Println("sum", sum)

	fmt.Println("------------------")

	sum = 1
	for sum < 1000 {
		sum += sum // sum = sum * 2

	}
	fmt.Println("sum", sum)
}
