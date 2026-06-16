package main

import "fmt"

func average(score []float64) float64 {
	if len(score) == 0 {
		return 0.0
	}

	sum := 0.0
	for _, s := range score {
		sum += s
	}

	return sum / float64(len(score))
}

func main() {

	r := average([]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(r)

}
