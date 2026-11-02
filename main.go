package main

import "fmt"

func main() {

	average := calculateAverage(10, 20, 30, 40)

	fmt.Println("average:", average)
}

func calculateAverage(numbers ...int) float64 {
	total := 0

	if len(numbers) == 0 {
		return float64(total)
	}

	for _, number := range numbers {
		total += number
	}

	average := float64(total) / float64(len(numbers))

	return float64(average)
}
