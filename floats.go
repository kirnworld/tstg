package main

import (
	"bufio"
	"os"
	"strconv"
)

// Reads float54 data from each line of a file
func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := os.Open("data.txt")
	if err != nil {
		return numbers, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		numbers = append(numbers, number)
	}
	err = file.Close()
	if err != nil {
		return numbers, err
	}
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
	return numbers, nil
}
