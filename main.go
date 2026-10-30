package main

import "fmt"

func main() {

	numbers := [5]int{10, 20, 30, 40, 50}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	fmt.Println(numbers[0])
	fmt.Println(numbers[4])
}
