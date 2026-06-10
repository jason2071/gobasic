package main

import (
	"errors"
	"fmt"
	"math"
)

func FunctionValue() {
	a, _ := apply(1, 2, add)
	b, _ := apply(1, 2, sub)
	c, _ := apply(1, 2, nil)
	fmt.Println(a, b, c)
}

func apply(a, b float64, op func(float64, float64) float64) (float64, error) {
	if op == nil {
		return math.NaN(), errors.New("apply: nil operation")
	}
	return op(a, b), nil
}

func add(a, b float64) float64 {
	return a + b
}

func sub(a, b float64) float64 {
	return a - b
}
