package main

import "fmt"

func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isPrime(7))   // true
	fmt.Println(isPrime(100)) // false
	fmt.Println(isPrime(1))   // false (1 ไม่ใช่จำนวนเฉพาะ)
}
