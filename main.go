package main

import "fmt"

func main() {
    numbers := []int{10, 20, 30}

    numbers = append(numbers, 40)
	numbers = append(numbers, 50)

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

    fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
}