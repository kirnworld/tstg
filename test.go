package main

import (
	"fmt"
	"math"
)

const (
	first = iota
	second
	third
	fourth
	fifth
	sixth
)

func main() {
	fmt.Println(third)
	var x float64 = 5
	var n float64 = 2
	var v = math.Pow(x, n)
	switch v {
	case 0:
		fmt.Println("Zero")
	default:
		fmt.Println("def")
	}

	sum := 0
	for i := 1; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	var h int
	for fmt.Scan(&h); h != 0; fmt.Scan(&h) {
		fmt.Println(h)
	}

}
