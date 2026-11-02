package main

import "fmt"

func main() {

	result := calculateArea(10, 5)

	fmt.Println("Area:", result)

}

func calculateArea(width int, height int) int {

	return width * height
}
