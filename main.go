package main

import "fmt"

func main() {
	age := 20

	if age >= 60 {
		fmt.Println("Senior")
	} else if age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Minor")
	}
}
