package main

import (
	"fmt"
	"math"
)

func maximum(numbers ...float64) float64 {
	max := math.Inf(-1)
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}

func inRange(mn float64, mx float64, cells ...float64) []float64 {
	var result []float64
	for _, cell := range cells {
		if cell >= mn && cell <= mx {
			result = append(result, cell)
		}
	}
	return result
}

func main() {
	fmt.Println(maximum(2.58, 1.98, 3.67, 54.2, 27.6))
	fmt.Println(inRange(2.6, 40.5, 5.2, 8.9, 4.2, 80.5, 60.1, 21.3))
}
