package main

import "fmt"

func main() {

	x := 8 // binary: 1000

	//ย้ายบิตไปทางซ้าย n ตำแหน่ง (คูณด้วย 2^n)
	y := x << 1    // 10000 = 16 (8 × 2^1)
	fmt.Println(y) // 16

	//ย้ายบิตไปทางขวา n ตำแหน่ง (หารด้วย 2^n)
	z := x >> 1    // 100 = 4 (8 ÷ 2^1)
	fmt.Println(z) // 4
}
