package main

import (
	"fmt"
	"log"
)

var metersPerLiter float64

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("invalid width %0.2f", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("invalid height %0.2f", height)
	}
	area := width * height
	return area / 10.0, nil
}

func double(number *int) {
	fmt.Println(*number)
	*number *= 2
}

func main() {
	amount2 := 6
	double(&amount2)
	fmt.Println(amount2)

	amount, err := paintNeeded(4.2, -3.2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2f liters needed\n", amount)

}
