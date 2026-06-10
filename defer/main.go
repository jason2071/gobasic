package main

import "fmt"

func main() {

	defer fmt.Println("john")
	fmt.Println("hello")

	fmt.Println("counting")

	for i := range 10 {
		defer fmt.Println(i) // 0,1,2,3,4,5,6,7,8,9
	}

	fmt.Println("done")
}
