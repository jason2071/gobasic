package numbers

import (
	"fmt"
	"math"
)

func NumberAvg() {
	fmt.Println("int8:")
	fmt.Printf("MinInt8: %d\n", math.MinInt8)
	fmt.Printf("MaxInt8: %d\n", math.MaxInt8)

	fmt.Println("float32:")
	fmt.Printf("SmallestNonzeroFloat32: %g\n", math.SmallestNonzeroFloat32)
	fmt.Printf("MaxFloat32: %g\n", math.MaxFloat32)

	resp := avg(1, 2)
	fmt.Printf("result: %.2f\n", resp)

	numbers := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	resp2 := avg2(numbers)
	fmt.Printf("result: %.2f\n", resp2)
}

func avg(a float64, b float64) float64 {
	return (a + b) / 2
}

func avg2(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}

	var sum float64
	for _, v := range data {
		sum += v
	}

	return sum / float64(len(data))
}
