package main

import (
	"fmt"
	"strings"
)

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func main() {
	// primes := [6]int{2, 3, 5, 7, 11, 13}

	// var s []int = primes[1:4]
	// fmt.Println(s)

	// var s2 []int = primes[1:]
	// fmt.Println(s2)

	// var s3 []int = primes[:4]
	// fmt.Println(s3)

	// var s4 []int = primes[:]
	// fmt.Println(s4)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "xxx"
	fmt.Println(a, b)

	fmt.Println(names)

	fmt.Println("------------------")

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	fmt.Println("------------------")

	a1 := make([]int, 5)
	printSlice("a", a1)

	// b1 := make([]int, 0, 5)
	// printSlice("b", b1)

	// c1 := b1[:2]
	// printSlice("c", c1)

	// d1 := c1[2:5]
	// printSlice("d", d1)

	a1 = append(a1, 0)
	printSlice("a", a1)
	a1 = append(a1, 0)
	printSlice("a", a1)
	a1 = append(a1, 0)
	printSlice("a", a1)
	a1 = append(a1, 0)
	printSlice("a", a1)
	a1 = append(a1, 0)
	printSlice("a", a1)
	a1 = append(a1, 0)
	printSlice("a", a1)

	fmt.Println("------------------")

	// Tic Tac Toe XOXO
	// Create a tic-tac-toe board.
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	// board[2][2] = "O"
	// board[1][2] = "X"
	// board[1][0] = "O"
	// board[0][2] = "X"
	// board[0][1] = "O"
	// board[1][1] = "X"
	// board[2][0] = "O"
	// board[2][1] = "X"

	for i := range board {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

}
