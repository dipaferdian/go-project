package main

import "fmt"

func main() {

	sum, substactrion, multilpy := calculate(10, 5)

	fmt.Println("sum:", sum)
	fmt.Println("substactrion:", substactrion)
	fmt.Println("multilpy:", multilpy)

}

func calculate(a int, b int) (int, int, int) {

	return a + b, a - b, a * b
}
